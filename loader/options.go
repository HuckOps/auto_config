package loader

import (
	"context"
	"github.com/huckops/auto_config/reader"
	"github.com/huckops/auto_config/source"
)

type Options struct {
	Source  []source.Source
	Context context.Context
	Reader  reader.Reader
}
type Option func(options *Options)

func WithSource(s source.Source) Option {
	return func(options *Options) {
		options.Source = append(options.Source, s)
	}
}

func WithReader(r reader.Reader) Option {
	return func(options *Options) {
		options.Reader = r
	}
}
