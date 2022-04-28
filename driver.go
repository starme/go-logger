package logger

import (
	"github.com/starme/logger/handles"
	"go.uber.org/zap"
)

type Driver struct {
	handler handles.Handler
}

func MakeHandle(config Channel) (handles.Handler, bool) {
	switch config.Type {
	case "simple":
		return handles.MakeSimpleHandle(config.Path, config.Level), true
	case "daily":
		return handles.MakeDailyHandle(config.Path, config.Level, config.Days), true
	}
	return nil, false
}

func (d Driver) Info(msg string, fields ...zap.Field) {
	d.handler.Info(msg, fields...)
}