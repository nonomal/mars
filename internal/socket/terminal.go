package socket

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/duc-cnzj/mars/api/v5/types"
	websocket_pb "github.com/duc-cnzj/mars/api/v5/websocket"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/repo"
	"github.com/duc-cnzj/mars/v5/internal/util/closeable"
	"k8s.io/client-go/tools/remotecommand"
)

var (
	ETX                 = []byte("\u0003")
	END_OF_TRANSMISSION = []byte("\u0004")
)

const (
	OpStdout = "stdout"
	OpStdin  = "stdin"
	OpResize = "resize"
	OpToast  = "toast"
)

type sizeStore struct {
	rwMu          sync.RWMutex
	width, height uint16
	reset         bool
}

func (s *sizeStore) ResetTerminalRowCol(reset bool) {
	s.rwMu.Lock()
	defer s.rwMu.Unlock()
	s.reset = reset
}

func (s *sizeStore) TerminalRowColNeedReset() bool {
	s.rwMu.RLock()
	defer s.rwMu.RUnlock()
	return s.reset
}

func (s *sizeStore) Set(width, height uint16) {
	s.rwMu.Lock()
	defer s.rwMu.Unlock()
	s.height = height
	s.width = width
}

func (s *sizeStore) Changed(width, height uint16) bool {
	s.rwMu.RLock()
	defer s.rwMu.RUnlock()
	if s.height != height {
		return true
	}
	if s.width != width {
		return true
	}

	return false
}

func (s *sizeStore) Width() uint16 {
	s.rwMu.RLock()
	defer s.rwMu.RUnlock()
	return s.width
}

func (s *sizeStore) Height() uint16 {
	s.rwMu.RLock()
	defer s.rwMu.RUnlock()
	return s.height
}

type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue

	Container() *repo.Container
	SetShell(string)
	Toast(string) error

	Send(ctx context.Context, message *websocket_pb.TerminalMessage) error
	Resize(remotecommand.TerminalSize) error

	Recorder() repo.Recorder

	ResetTerminalRowCol(bool)
	Height() uint16
	Width() uint16

	Close(context.Context, string) bool
	IsClosed() bool
}

type myPtyHandler struct {
	logger    mlog.Logger
	sessionID string
	container *repo.Container
	recorder  repo.Recorder
	eventRepo repo.EventRepo
	conn      Conn

	doneChan  chan struct{}
	sizeStore *sizeStore

	shellMu sync.RWMutex
	shellCh chan *websocket_pb.TerminalMessage

	sizeMu   sync.RWMutex
	sizeChan chan remotecommand.TerminalSize

	closeable.Closeable
}

func (t *myPtyHandler) SetShell(shell string) {
	t.recorder.SetShell(shell)
}

func (t *myPtyHandler) Container() *repo.Container {
	return t.container
}

func (t *myPtyHandler) Height() uint16 {
	return t.sizeStore.Height()
}

func (t *myPtyHandler) Width() uint16 {
	return t.sizeStore.Width()
}

func (t *myPtyHandler) Read(p []byte) (n int, err error) {
	var (
		msg *websocket_pb.TerminalMessage
		ok  bool
	)
	select {
	case <-t.doneChan:
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("[Websocket]: %v doneChan closed", t.sessionID)
	case msg, ok = <-t.shellCh:
		if !ok {
			return copy(p, END_OF_TRANSMISSION), fmt.Errorf("[Websocket]: %v channel closed", t.sessionID)
		}
	}

	switch msg.Op {
	case OpStdin:
		return copy(p, msg.Data), nil
	case OpResize:
		t.logger.Debugf("[Websocket]: resize width: %v  height: %v", msg.Width, msg.Height)
		t.Resize(remotecommand.TerminalSize{Width: uint16(msg.Width), Height: uint16(msg.Height)})
		return 0, nil
	default:
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Op)
	}
}

func (t *myPtyHandler) Write(p []byte) (n int, err error) {
	select {
	case <-t.doneChan:
		return len(p), fmt.Errorf("[Websocket]: %v doneChan closed", t.sessionID)
	default:
	}
	if t.IsClosed() {
		return len(p), fmt.Errorf("[Websocket]: %v ws already closed", t.sessionID)
	}

	if _, err = t.recorder.Write(p); err != nil {
		t.logger.Debugf("[Websocket]: %v recorder write failed: %v", t.sessionID, err)
	}
	if t.sizeStore.TerminalRowColNeedReset() && t.sizeStore.Width() != 0 {
		t.logger.Debugf("reset shell size height: %d, width: %d", t.sizeStore.Height(), t.sizeStore.Width())
		t.sizeStore.ResetTerminalRowCol(false)
		if err = t.Resize(remotecommand.TerminalSize{Width: t.sizeStore.Width(), Height: t.sizeStore.Height()}); err != nil {
			t.logger.Debugf("resize shell size failed: %v", err)
		}
	}
	NewMessageSender(t.conn, t.sessionID, WsHandleExecShellMsg).SendProtoMsg(&websocket_pb.WsHandleShellResponse{
		Metadata: &websocket_pb.Metadata{
			Id:     t.conn.ID(),
			Uid:    t.conn.UID(),
			Slug:   t.sessionID,
			Type:   WsHandleExecShellMsg,
			Result: ResultSuccess,
		},
		TerminalMessage: &websocket_pb.TerminalMessage{
			Op:        OpStdout,
			Data:      p,
			SessionId: t.sessionID,
		},
		Container: &websocket_pb.Container{
			Namespace: t.Container().Namespace,
			Pod:       t.Container().Pod,
			Container: t.Container().Container,
		},
	})

	return len(p), nil
}

func (t *myPtyHandler) ResetTerminalRowCol(reset bool) {
	t.sizeStore.ResetTerminalRowCol(reset)
}

func (t *myPtyHandler) Recorder() repo.Recorder {
	return t.recorder
}

func (t *myPtyHandler) Next() *remotecommand.TerminalSize {
	select {
	case size, ok := <-t.sizeChan:
		if !ok {
			return nil
		}
		if size.Width != 0 && size.Height != 0 {
			if t.sizeStore.Changed(size.Width, size.Height) {
				t.recorder.Resize(size.Width, size.Height)
			}
			t.sizeStore.Set(size.Width, size.Height)
		}
		return &size
	case <-t.doneChan:
		return nil
	}
}

func (t *myPtyHandler) Send(ctx context.Context, m *websocket_pb.TerminalMessage) error {
	t.shellMu.Lock()
	defer t.shellMu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.doneChan:
		close(t.shellCh)
		return errors.New("doneChan closed")
	default:
	}

	select {
	case t.shellCh <- m:
	default:
		t.logger.Warning("[Websocket]: shellCh chan full")
	}
	return nil
}

func (t *myPtyHandler) Resize(size remotecommand.TerminalSize) error {
	select {
	case <-t.doneChan:
		close(t.sizeChan)
		return errors.New("doneChan closed")
	default:
	}

	t.sizeMu.Lock()
	defer t.sizeMu.Unlock()
	select {
	case t.sizeChan <- size:
	default:
		return errors.New("sizeChan chan full")
	}
	return nil
}

func (t *myPtyHandler) IsClosed() bool {
	return t.Closeable.IsClosed()
}

func (t *myPtyHandler) CloseDoneChan() bool {
	if !t.Closeable.Close() {
		return false
	}
	close(t.doneChan)
	return true
}

func (t *myPtyHandler) Close(ctx context.Context, reason string) bool {
	if !t.Closeable.Close() {
		return false
	}
	NewMessageSender(t.conn, t.sessionID, WsHandleCloseShell).SendProtoMsg(&websocket_pb.WsHandleShellResponse{
		Metadata: &websocket_pb.Metadata{
			Id:     t.conn.ID(),
			Uid:    t.conn.UID(),
			Slug:   t.sessionID,
			Type:   WsHandleCloseShell,
			Result: ResultSuccess,
		},
		TerminalMessage: &websocket_pb.TerminalMessage{
			SessionId: t.sessionID,
			Op:        OpStdout,
			Data:      []byte(reason),
		},
		Container: &websocket_pb.Container{
			Namespace: t.Container().Namespace,
			Pod:       t.Container().Pod,
			Container: t.Container().Container,
		},
	})

	t.Send(ctx, &websocket_pb.TerminalMessage{
		Op:        OpStdin,
		Data:      ETX,
		SessionId: t.sessionID,
	})
	time.Sleep(200 * time.Millisecond)
	t.Send(ctx, &websocket_pb.TerminalMessage{
		Op:        OpStdin,
		Data:      END_OF_TRANSMISSION,
		SessionId: t.sessionID,
	})
	t.logger.Debug("[Websocket]: close shell.")
	if err := t.Recorder().Close(); err != nil {
		t.logger.Error(err)
	}
	recoder := t.Recorder()
	var fid int
	rf := recoder.File()
	if rf != nil {
		fid = rf.ID
	}
	t.eventRepo.FileAuditLogWithDuration(
		types.EventActionType_Shell,
		recoder.User().Name,
		fmt.Sprintf("用户进入容器执行命令，container: '%s', namespace: '%s', pod： '%s'", recoder.Container().Container, recoder.Container().Namespace, recoder.Container().Pod),
		fid,
		recoder.Duration(),
	)
	close(t.doneChan)
	return true
}

// Toast can be used to send the user any OOB messages
// hterm puts these in the center of the terminal
func (t *myPtyHandler) Toast(p string) error {
	NewMessageSender(t.conn, t.sessionID, WsHandleExecShellMsg).SendProtoMsg(&websocket_pb.WsHandleShellResponse{
		Metadata: &websocket_pb.Metadata{
			Id:     t.conn.ID(),
			Uid:    t.conn.UID(),
			Slug:   t.sessionID,
			Type:   WsHandleExecShellMsg,
			Result: ResultSuccess,
		},
		TerminalMessage: &websocket_pb.TerminalMessage{
			Op:        OpToast,
			Data:      []byte(p),
			SessionId: t.sessionID,
		},
		Container: &websocket_pb.Container{
			Container: t.Container().Container,
			Namespace: t.Container().Namespace,
			Pod:       t.Container().Pod,
		},
	})
	return nil
}

type SessionMapper interface {
	Get(sessionId string) (PtyHandler, bool)
	Set(sessionId string, session PtyHandler)
	CloseAll(ctx context.Context)
	Close(ctx context.Context, sessionId string, status uint32, reason string)
}

// sessionMap stores a map of all myPtyHandler objects and a sessLock to avoid concurrent conflict
type sessionMap struct {
	wg     sync.WaitGroup
	logger mlog.Logger

	sessLock sync.RWMutex
	Sessions map[string]PtyHandler
}

func NewSessionMap(logger mlog.Logger) SessionMapper {
	return &sessionMap{Sessions: map[string]PtyHandler{}, logger: logger}
}

// Get return a given terminalSession by sessionId
func (sm *sessionMap) Get(sessionId string) (PtyHandler, bool) {
	sm.sessLock.RLock()
	defer sm.sessLock.RUnlock()
	h, ok := sm.Sessions[sessionId]
	return h, ok
}

// Set store a myPtyHandler to sessionMap
func (sm *sessionMap) Set(sessionId string, session PtyHandler) {
	sm.sessLock.Lock()
	defer sm.sessLock.Unlock()
	sm.Sessions[sessionId] = session
}

func (sm *sessionMap) CloseAll(ctx context.Context) {
	sm.logger.Debug("[Websocket]: close all.")
	sm.sessLock.Lock()
	defer sm.sessLock.Unlock()

	for _, s := range sm.Sessions {
		sm.wg.Add(1)
		go func(s PtyHandler) {
			defer sm.wg.Done()
			s.Close(ctx, "websocket conn closed")
		}(s)
	}
	sm.wg.Wait()
	sm.Sessions = map[string]PtyHandler{}
}

// Close shuts down the SockJS connection and sends the status code and reason to the client
// Can happen if the process exits or if there is an error starting up the process
// For now the status code is unused and reason is shown to the user (unless "")
func (sm *sessionMap) Close(ctx context.Context, sessionId string, status uint32, reason string) {
	sm.logger.Debugf("[Websocket]: session %v closed, reason: %s, status: %v.", sessionId, reason, status)
	sm.sessLock.Lock()
	defer sm.sessLock.Unlock()
	if s, ok := sm.Sessions[sessionId]; ok {
		delete(sm.Sessions, sessionId)
		sm.wg.Add(1)
		go func() {
			defer sm.wg.Done()
			s.Close(ctx, reason)
		}()
	}
}

// startProcess is called by handleAttach
// Executed cmd in the contracts.Container specified in request and connects it up with the ptyHandler (a session)
func (wc *WebsocketManager) startProcess(ctx context.Context, container *repo.Container, cmd []string, ptyHandler PtyHandler) error {
	return wc.k8sRepo.Execute(ctx, container, &repo.ExecuteInput{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		TTY:               true,
		Cmd:               cmd,
		TerminalSizeQueue: ptyHandler,
	})
}

// isValidShell checks if the shell is an allowed one
func isValidShell(validShells []string, shell string) bool {
	for _, validShell := range validShells {
		if validShell == shell {
			return true
		}
	}
	return false
}

var silenceShellExitMessages = []string{
	"command terminated with exit code 126",
	"command terminated with exit code 130",
}

func silence(err error) bool {
	for _, message := range silenceShellExitMessages {
		if strings.Contains(err.Error(), message) {
			return true
		}
	}
	return false
}

// WaitForTerminal is called from apihandler.handleAttach as a goroutine
// Waits for the SockJS connection to be opened by the client the session to be bound in handleMyPtyHandler
func (wc *WebsocketManager) WaitForTerminal(ctx context.Context, conn Conn, container *repo.Container, shell, sessionId string) {
	defer func() {
		wc.logger.Debugf("[Websocket]: WaitForTerminal EXIT: total go: %v", runtime.NumGoroutine())
	}()
	var err error
	validShells := []string{"bash", "sh", "powershell", "cmd"}
	session, got := conn.GetPtyHandler(sessionId)
	if !got {
		return
	}
	if isValidShell(validShells, shell) {
		cmd := []string{shell}
		session.SetShell(shell)
		err = wc.startProcess(ctx, container, cmd, session)
	} else {
		// No shell given or it was not valid: try some shells until one succeeds or all fail
		// FIXME: if the first shell fails then the first keyboard event is lost
		for idx, testShell := range validShells {
			wc.logger.Debug("try: " + testShell)
			if session.IsClosed() {
				wc.logger.Debugf("session 已关闭，不会继续尝试连接其他 shell: '%s'", strings.Join(validShells[idx:], ", "))
				break
			}
			cmd := []string{testShell}
			session.SetShell(testShell)
			if err = wc.startProcess(ctx, container, cmd, session); err == nil {
				break
			}
			// 当出现 bash 回退的时候，需要注意，resize 不会触发，导致，新的 'sh', width, height 和用户端不一致，所以需要重置，
			// 通过 sizeStore 记录上次用户的 height, width, 当 bash 回退时，在用户输入时应用到新的 sh 中
			session = wc.resetSession(session)
			conn.SetPtyHandler(sessionId, session)
		}
	}

	if err != nil {
		wc.logger.Debugf("[Websocket]: %v", err.Error())
		if !silence(err) {
			session.Toast(err.Error())
		}
		conn.ClosePty(ctx, sessionId, 2, err.Error())
		return
	}

	conn.ClosePty(ctx, sessionId, 1, "Process exited")
}

func (wc *WebsocketManager) resetSession(session PtyHandler) PtyHandler {
	var width, height uint16 = 106, 25
	func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		af := time.NewTimer(3 * time.Second)
		defer ticker.Stop()
		defer af.Stop()
		for session.Width() == 0 {
			select {
			case <-ticker.C:
				wc.logger.Debug("sleep....")
				break
			case <-af.C:
				wc.logger.Warningf("can't get previous width,height, use default height: 25, width: 106.")
				return
			}
		}
		width = session.Width()
		height = session.Height()
	}()
	wc.logger.Debug("done....")

	spty := session.(*myPtyHandler)
	var newSession PtyHandler = session
	if spty.CloseDoneChan() {
		newSession = &myPtyHandler{
			logger:    spty.logger,
			sessionID: spty.sessionID,
			container: spty.container,
			recorder:  spty.recorder,
			eventRepo: spty.eventRepo,
			conn:      spty.conn,
			doneChan:  make(chan struct{}),
			sizeChan:  make(chan remotecommand.TerminalSize, 1),
			shellCh:   make(chan *websocket_pb.TerminalMessage, 500),
			sizeStore: &sizeStore{
				width:  width,
				height: height,
				reset:  true,
			},
		}
	}
	return newSession
}

type TerminalResponse struct {
	ID string `json:"sessionID"`
}

func checkSessionID(container *websocket_pb.Container, id string) bool {
	prefix := fmt.Sprintf("%s-%s-%s:", container.Namespace, container.Pod, container.Container)
	return strings.HasPrefix(id, prefix)
}

func (wc *WebsocketManager) StartShell(ctx context.Context, input *websocket_pb.WsHandleExecShellInput, conn Conn) (string, error) {
	var (
		container = &repo.Container{
			Namespace: input.Container.Namespace,
			Pod:       input.Container.Pod,
			Container: input.Container.Container,
		}
		sessionID = input.SessionId
	)

	if !checkSessionID(input.Container, sessionID) {
		return "", fmt.Errorf("invalid session sessionID, must format: '<namespace>-<pod>-<container>:<randomID>', input: '%s'", sessionID)
	}

	pty := &myPtyHandler{
		logger:    wc.logger,
		sessionID: sessionID,
		eventRepo: wc.eventRepo,
		container: container,
		recorder:  wc.fileRepo.NewRecorder(conn.GetUser(), container),
		conn:      conn,
		doneChan:  make(chan struct{}),
		sizeStore: &sizeStore{},
		shellCh:   make(chan *websocket_pb.TerminalMessage, 500),
		sizeChan:  make(chan remotecommand.TerminalSize, 1),
	}
	conn.SetPtyHandler(sessionID, pty)

	go func() {
		defer wc.logger.HandlePanic("Websocket: WaitForTerminal")
		wc.WaitForTerminal(ctx, conn, container, "", sessionID)
	}()

	return sessionID, nil
}
