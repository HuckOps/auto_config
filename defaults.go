package auto_config

import (
	"auto_config/loader"
	"auto_config/reader"
)

type Config struct {
	opts     Options
	snapshot loader.SnapShot
}

func NewConfig(opts ...Option) (Config, error) {
	options := Options{
		Reader: reader.NewReader(),
	}
	for _, opt := range opts {
		opt(&options)
	}

	// 源文件放入装载器
	options.Loader = loader.NewLoader(loader.WithReader(options.Reader))
	// 首次读取文件
	err := options.Loader.Load(options.Source...)
	if err != nil {
		return Config{}, err
	}
	config := Config{
		opts: options,
	}
	config.Init()
	return config, nil
}

func (config *Config) Init() {
	// 获取初始状态
	config.snapshot = config.opts.Loader.Snapshot()
	//
}
