package handles

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type SimpleHandle struct {
	path	string
	level	string

	writer	*zap.Logger
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
		encoder(),
		zapcore.WriteSyncer(file),
		level(handle.level),
	)
	handle.writer = zap.New(core, zap.AddCaller())
}

func (handle SimpleHandle) Debug(msg string, fields ...zap.Field) {
	handle.writer.Debug(msg, fields...)
}

func (handle SimpleHandle) Info(msg string, fields ...zap.Field) {
	handle.writer.Info(msg, fields...)
}

func (handle SimpleHandle) Infof(msg string, fields ...interface{}) {
	handle.writer.Sugar().Infof(msg, fields...)
}

