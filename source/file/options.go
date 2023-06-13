package file

import (
	"context"
	"github.com/huckops/auto_config/source"
)

//type Options struct {
//}
type fileKeyPath struct {
}

func WithPath(path string) source.Option {
	return func(options *source.Options) {
		options.Context = context.WithValue(options.Context, fileKeyPath{}, path)
	}
}
