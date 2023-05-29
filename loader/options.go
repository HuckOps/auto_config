package loader

import (
	"auto_config/reader"
	"auto_config/source"
	"context"
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
