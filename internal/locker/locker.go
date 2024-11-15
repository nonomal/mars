package locker

import (
	"github.com/duc-cnzj/mars/v5/internal/config"
	"github.com/duc-cnzj/mars/v5/internal/data"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/util/timer"
)

type Locker interface {
	ID() string
	Type() string
	Acquire(key string, seconds int64) bool
	RenewalAcquire(key string, seconds int64, renewalSeconds int64) (releaseFn func(), acquired bool)
	Release(key string) bool
	ForceRelease(key string) bool
	Owner(key string) string
}

func NewLocker(cfg *config.Config, data data.Data, logger mlog.Logger, timer timer.Timer) (Locker, error) {
	logger = logger.WithModule("locker/locker")
	if cfg.DBDriver == "sqlite" && cfg.CacheDriver == "db" {
		cfg.CacheDriver = "memory"
		logger.Warning(`使用 DBDriver 为 "sqlite" 时，CacheDriver,Locker 只能使用 "memory"!`)
	}
	driver := cfg.CacheDriver

	switch driver {
	case "db":
		return NewDatabaseLock(timer, [2]int{2, 100}, data, logger), nil
	case "memory":
		fallthrough
	default:
		return NewMemoryLock(timer, [2]int{2, 100}, NewMemStore(), logger), nil
	}
}
