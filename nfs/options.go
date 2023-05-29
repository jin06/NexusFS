package nfs

import "github.com/sirupsen/logrus"

type Option func(*Options)

type FSTYPE string

const (
	FSTYPE_DEFAULT   = ""
	FSTYYPE_LOOKBACK = "lookback"
	FSTYPE_DB        = "db"
)

type Options struct {
	MountPoint   string
	FSType       FSTYPE
	LookBackPath string
	FuseName string
}

func NewOptions(opt ...Option) Options {
	opts := Options{}
	for _, o := range opt {
		o(&opts)
	}
	logrus.Debugf("%v", opts)
	return opts
}

func MountPoint(m string) Option {
	return func(options *Options) {
		options.MountPoint = m
	}
}

func FSType(s string) Option {
	return func(options *Options) {
		options.FSType = FSTYPE(s)
	}
}

func LBPath(s string) Option {
	return func(options *Options) {
		options.LookBackPath = s
	}
}

func FuseName(s string) Option {
	return func(options *Options) {
		options.FuseName = s
	}
}
