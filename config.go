package auto_config

import (
	"github.com/huckops/auto_config/loader"
	"github.com/huckops/auto_config/reader"
	"github.com/huckops/auto_config/source"
)

type Options struct {
	Reader       reader.Reader
	Entity       interface{}
	Loader       loader.Loader
	Source       []source.Source
	InitCallBack []InitCallback
}

type Option func(options *Options)
type InitCallback func()
