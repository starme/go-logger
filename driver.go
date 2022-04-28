package logger

import (
	"github.com/starme/logger/handles"
	"go.uber.org/zap"
)

func MakeHandle(config Channel) (handles.Handler, bool) {
	switch config.Type {
	case "simple":
		return handles.MakeSimpleHandle(config.Path, config.Level), true
	case "daily":
		return handles.MakeDailyHandle(config.Path, config.Level, config.Days), true
	}
	return nil, false
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
