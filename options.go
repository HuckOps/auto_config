package auto_config

import "auto_config/source"

func WithEntity(v interface{}) Option {
	return func(options *Options) {
		options.Entity = v
	}
}

func WithSource(source source.Source) Option {
	return func(options *Options) {
		options.Source = append(options.Source, source)
	}
}
