package handles

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type StdHandler struct {
	writer *zap.Logger
}

func MakeStdHandle() *StdHandler {
	handler := &StdHandler{}
	handler.init()
	return handler
}

func (handle *StdHandler) init() {
	stdout := zapcore.NewCore(
		consoleEncoder(),
		zapcore.Lock(os.Stdout),
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l <= zap.WarnLevel
		}),
	)

	stderr := zapcore.NewCore(
		consoleEncoder(),
		zapcore.Lock(os.Stderr),
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l > zap.WarnLevel
		}),
	)
	handle.writer = zap.New(zapcore.NewTee(stdout, stderr), zap.AddCaller(), zap.Development())
}

func (handle StdHandler) Debug(msg string, fields ...zap.Field) {
	handle.writer.Debug(msg, fields...)
}

func (handle StdHandler) Debugf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Debugf(msg, fields...)
}

func (handle StdHandler) Info(msg string, fields ...zap.Field) {
	handle.writer.Info(msg, fields...)
}

func (handle StdHandler) Infof(msg string, fields ...interface{}) {
	handle.writer.Sugar().Infof(msg, fields...)
}

func (handle StdHandler) Warn(msg string, fields ...zap.Field) {
	handle.writer.Warn(msg, fields...)
}

func (handle StdHandler) Warnf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Warnf(msg, fields...)
}

func (handle StdHandler) Error(msg string, fields ...zap.Field) {
	handle.writer.Error(msg, fields...)
}

func (handle StdHandler) Errorf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Errorf(msg, fields...)
}

func (handle StdHandler) DPanic(msg string, fields ...zap.Field) {
	handle.writer.DPanic(msg, fields...)
}

func (handle StdHandler) DPanicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().DPanicf(msg, fields...)
}

func (handle StdHandler) Panic(msg string, fields ...zap.Field) {
	handle.writer.Panic(msg, fields...)
}

func (handle StdHandler) Panicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Panicf(msg, fields...)
}

func (handle StdHandler) Fatal(msg string, fields ...zap.Field) {
	handle.writer.Fatal(msg, fields...)
}

func (handle StdHandler) Fatalf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Fatalf(msg, fields...)
}
