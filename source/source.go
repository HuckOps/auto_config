package source

import (
	"time"
)

// 存放数据
type ChangeSet struct {
	Data     []byte
	CheckSum string
	Format   string
	//Source    string
	Timestamp time.Time
}

type Source interface {
	Read() (*ChangeSet, error)
	Write(*ChangeSet) (*ChangeSet, error)
	Watcher() (Watcher, error)
	Path() string
}

type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}
