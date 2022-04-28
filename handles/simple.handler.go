package handles

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type SimpleHandle struct {
	path  string
	level string

	writer *zap.Logger
}

func MakeSimpleHandle(path string, level string) *SimpleHandle {
	handler := &SimpleHandle{
		path: path, level: level,
	}
	handler.init()
	return handler
}

func (handle *SimpleHandle) init() {
	file, _ := os.Create(handle.path)
	core := zapcore.NewCore(
		jsonEncoder(),
		zapcore.WriteSyncer(file),
		level(handle.level),
	)
	handle.writer = zap.New(core, zap.AddCaller())
}

func (handle SimpleHandle) Debug(msg string, fields ...zap.Field) {
	handle.writer.Debug(msg, fields...)
}

func (handle SimpleHandle) Debugf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Debugf(msg, fields...)
}

func (handle SimpleHandle) Info(msg string, fields ...zap.Field) {
	handle.writer.Info(msg, fields...)
}

func (handle SimpleHandle) Infof(msg string, fields ...interface{}) {
	handle.writer.Sugar().Infof(msg, fields...)
}

func (handle SimpleHandle) Warn(msg string, fields ...zap.Field) {
	handle.writer.Warn(msg, fields...)
}

func (handle SimpleHandle) Warnf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Warnf(msg, fields...)
}

func (handle SimpleHandle) Error(msg string, fields ...zap.Field) {
	handle.writer.Error(msg, fields...)
}

func (handle SimpleHandle) Errorf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Errorf(msg, fields...)
}

func (handle SimpleHandle) DPanic(msg string, fields ...zap.Field) {
	handle.writer.DPanic(msg, fields...)
}

func (handle SimpleHandle) DPanicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().DPanicf(msg, fields...)
}

func (handle SimpleHandle) Panic(msg string, fields ...zap.Field) {
	handle.writer.Panic(msg, fields...)
}

func (handle SimpleHandle) Panicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Panicf(msg, fields...)
}

func (handle SimpleHandle) Fatal(msg string, fields ...zap.Field) {
	handle.writer.Fatal(msg, fields...)
}

func (handle SimpleHandle) Fatalf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Fatalf(msg, fields...)
}
