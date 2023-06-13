package loader

import (
	"github.com/huckops/auto_config/source"
	"time"
)

type Loader interface {
	Load(...source.Source) error
	Watcher() Watcher
	Snapshot() SnapShot
	EnableReaderPanicSkip()
}
type Watcher interface {
	//Update(snapshot SnapShot)
	Next() (*SnapShot, error)
	Stop() error
}
type SnapShot struct {
	Version    time.Time
	LastChange *source.ChangeSet
}
