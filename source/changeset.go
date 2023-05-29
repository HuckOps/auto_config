package source

import (
	"crypto/md5"
	"fmt"
)

func (cs *ChangeSet) Sum() string {
	h := md5.New()
	h.Write(cs.Data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
