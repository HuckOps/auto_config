package reader

import "auto_config/source"

type Reader interface {
	Merge(...*source.ChangeSet) (*source.ChangeSet, error)
	Scan([]byte, interface{}) error
}
