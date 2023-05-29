package loader

import (
	"auto_config/source"
	"time"
)

type Loader interface {
	Load(...source.Source) error

	Snapshot() SnapShot
}

type SnapShot struct {
	Version    time.Time
	LastChange *source.ChangeSet
}
