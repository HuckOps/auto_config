package auto_config

import (
	"auto_config/loader"
	"auto_config/reader"
	"auto_config/source"
)

type Options struct {
	Reader reader.Reader
	Entity interface{}
	Loader loader.Loader
	Source []source.Source
}

type Option func(options *Options)
