package source

import (
	"context"
	"github.com/huckops/auto_config/encoder"
	"github.com/huckops/auto_config/encoder/json"
)

type Options struct {
	Encoder encoder.Encoder
	Context context.Context
}

type Option func(*Options)

func WithEncoder(e encoder.Encoder) Option {
	return func(options *Options) {
		options.Encoder = e
	}
}

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoder: json.NewEncoder(),
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
