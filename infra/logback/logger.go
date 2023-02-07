package logback

import "github.com/vela-ssoc/manager/infra/conf"

type Logger interface {
	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

func New(cfg conf.Logger) Logger {
	return nil
}
