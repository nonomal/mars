package socket

//go:generate mockgen -destination ../mock/mock_recorder.go -package mock github.com/duc-cnzj/mars/internal/socket RecorderInterface

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/duc-cnzj/mars-client/v4/types"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/utils"
	"github.com/duc-cnzj/mars/internal/utils/date"
)

type RecorderInterface interface {
	Write(data string) (err error)
	Close() error
	SetShell(string)
	GetShell() string
}

type Recorder struct {
	sync.RWMutex
	filepath  string
	container Container
	f         contracts.File
	startTime time.Time

	t    *MyPtyHandler
	once sync.Once

	buffer *bufio.Writer

	shellMu sync.RWMutex
	shell   string
}

var (
	startLine = "{\"version\": 2, \"width\": 106, \"height\": 25, \"timestamp\": %d, \"env\": {\"SHELL\": \"%s\", \"TERM\": \"xterm-256color\"}}\n"
	writeLine = "[%.6f, \"o\", %s]\n"
)

func (r *Recorder) GetShell() string {
	r.shellMu.RLock()
	defer r.shellMu.RUnlock()
	return r.shell
}

func (r *Recorder) SetShell(sh string) {
	r.shellMu.Lock()
	defer r.shellMu.Unlock()
	r.shell = sh
}

func (r *Recorder) Write(data string) (err error) {
	r.Lock()
	defer r.Unlock()
	r.once.Do(func() {
		var file contracts.File
		file, err = app.Uploader().Disk("shell").NewFile(fmt.Sprintf("%s/%s/%s",
			r.t.conn.GetUser().Name,
			time.Now().Format("2006-01-02"),
			fmt.Sprintf("recorder-%s-%s-%s-%s.cast", r.t.Namespace, r.t.Pod, r.t.Container.Container, utils.RandomString(20))))
		if err != nil {
			return
		}
		r.f = file
		r.buffer = bufio.NewWriterSize(r.f, 1024*20)
		r.filepath = file.Name()
		r.startTime = time.Now()
		r.buffer.Write([]byte(fmt.Sprintf(startLine, r.startTime.Unix(), r.shell)))
	})
	if err != nil {
		return err
	}
	marshal, _ := json.Marshal(data)
	_, err = r.buffer.WriteString(fmt.Sprintf(writeLine, float64(time.Since(r.startTime).Microseconds())/1000000, string(marshal)))
	return err
}

func (r *Recorder) Close() error {
	r.RLock()
	defer r.RUnlock()
	var (
		err      error
		uploader = app.Uploader()
	)
	if r.buffer == nil || r.startTime.IsZero() {
		return nil
	}
	r.buffer.Flush()
	stat, _ := r.f.Stat()
	var emptyFile bool = true
	if stat.Size() > 0 {
		file := &models.File{
			UploadType: uploader.Type(),
			Path:       r.filepath,
			Size:       uint64(stat.Size()),
			Username:   r.t.conn.GetUser().Name,
			Namespace:  r.container.Namespace,
			Pod:        r.container.Pod,
			Container:  r.container.Container,
		}
		app.DB().Create(file)
		var emodal = models.Event{
			Action:   uint8(types.EventActionType_Shell),
			Username: r.t.conn.GetUser().Name,
			Message:  fmt.Sprintf("用户进入容器执行命令，container: '%s', namespace: '%s', pod： '%s'", r.container.Container, r.container.Namespace, r.container.Pod),
			FileID:   &file.ID,
			Duration: date.HumanDuration(time.Since(r.startTime)),
		}

		app.DB().Create(&emodal)
		emptyFile = false
	}
	err = r.f.Close()
	if emptyFile {
		uploader.Delete(r.f.Name())
	}
	return err
}
