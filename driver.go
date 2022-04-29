package logger

import (
	"github.com/starme/logger/handles"
	"go.uber.org/zap"
)

func MakeHandle(config Channel) (handles.Handler, bool) {
	var handler handles.Handler
	switch config.Type {
	case Simple:
		handler = handles.MakeSimpleHandle(config.Path, config.Level, config.CallerEnable)
	case Daily:
		handler = handles.MakeDailyHandle(config.Path, config.Level, config.MaxAge, config.CallerEnable)
	case Std:
		handler = handles.MakeStdHandle()
	}
	if handler == nil {
		return handler, false
	}
	return handler, true
}

func Debug(msg string, fields ...zap.Field) {
	With(defDriver).Debug(msg, fields...)
}

func Debugf(msg string, fields ...interface{}) {
	With(defDriver).Debugf(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	With(defDriver).Info(msg, fields...)
}

func Infof(msg string, fields ...interface{}) {
	With(defDriver).Infof(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	With(defDriver).Warn(msg, fields...)
}

func Warnf(msg string, fields ...interface{}) {
	With(defDriver).Warnf(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	With(defDriver).Error(msg, fields...)
}

func Errorf(msg string, fields ...interface{}) {
	With(defDriver).Errorf(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	With(defDriver).DPanic(msg, fields...)
}

func DPanicf(msg string, fields ...interface{}) {
	With(defDriver).DPanicf(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	With(defDriver).Panic(msg, fields...)
}

func Panicf(msg string, fields ...interface{}) {
	With(defDriver).Panicf(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	With(defDriver).Fatal(msg, fields...)
}

func Fatalf(msg string, fields ...interface{}) {
	With(defDriver).Fatalf(msg, fields...)
}
