package store

type Option func(options *Options)

type Options struct {
	BoltStorePath string
}

func NewOptions(options ...Option) Options {
	opts := Options{}
	for _, o := range options {
		o(&opts)
	}
	return opts
}

func BoltStoreVal(path string) Option {
	return func(options *Options) {
		options.BoltStorePath = path
	}
}
