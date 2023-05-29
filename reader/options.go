package reader

import (
	"auto_config/encoder"
	"auto_config/encoder/json"
	"auto_config/encoder/yaml"
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
