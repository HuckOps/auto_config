package reader

import "github.com/huckops/auto_config/source"

type Reader interface {
	Merge(...*source.ChangeSet) (*source.ChangeSet, error)
	Scan([]byte, interface{}) error
}
