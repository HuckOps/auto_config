package reader

import (
	"github.com/huckops/auto_config/encoder"
	"github.com/huckops/auto_config/encoder/json"
	"github.com/huckops/auto_config/encoder/yaml"
)

type Options struct {
	Encoding map[string]encoder.Encoder
}

type Option func(options *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"yaml": yaml.NewEncoder(),
			"json": json.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
