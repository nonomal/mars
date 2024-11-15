package socket

//go:generate mockgen -destination ./mock_socket.go -package socket github.com/duc-cnzj/mars/v5/internal/socket JobManager,Job,Percentable,Conn,PtyHandler,TaskManager,GorillaWs,SessionMapper,ReleaseInstaller,DeployMsger,SafeWriteMessageChan
import (
	"github.com/duc-cnzj/mars/v5/internal/util/counter"
	"github.com/google/wire"
)

var WireSocket = wire.NewSet(NewWebsocketManager, NewJobManager, counter.NewCounter, NewReleaseInstaller)
