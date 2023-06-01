package auto_config

import (
	"auto_config/loader"
	"auto_config/reader"
	"sync"
)

type Config struct {
	opts     Options
	snapshot loader.SnapShot
	sync.RWMutex
}

func NewConfig(opts ...Option) (*Config, error) {
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
		return &Config{}, err
	}
	config := Config{
		opts: options,
	}
	err = config.Init()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) Init() error {
	// 获取初始状态
	config.snapshot = config.opts.Loader.Snapshot()
	// 初次解析
	err := config.Scan()
	if err != nil {
		return err
	}
	// 初次运行初始化
	for _, callback := range config.opts.InitCallBack {
		callback()
	}
	return nil
}

func (config *Config) Scan() error {
	return config.opts.Reader.Scan(config.snapshot.LastChange.Data, &config.opts.Entity)
}

func (config *Config) Watcher() {
	watch := func(w loader.Watcher) {
		for {
			snapshot, err := w.Next()
			if err != nil {
				panic(err)
			}
			config.Lock()
			config.snapshot = *snapshot
			err = config.Scan()
			if err != nil {
				config.Unlock()
				panic(err)
			}

			for _, callback := range config.opts.InitCallBack {
				callback()
			}
			config.Unlock()
		}
	}
	loadWatcher := config.opts.Loader.Watcher()
	go watch(loadWatcher)
}
