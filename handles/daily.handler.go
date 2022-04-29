package handles

import (
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type DailyHandle struct {
	path  string
	level string
	daily int

	writer *zap.Logger
}

func MakeDailyHandle(path string, level string, daily int) *DailyHandle {
	handler := &DailyHandle{
		path: path, level: level, daily: daily,
	}
	handler.init()
	return handler
}

func (handle *DailyHandle) init() {
	core := zapcore.NewCore(
		jsonEncoder(),
		zapcore.AddSync(handle.syncer()),
		level(handle.level),
	)
	handle.writer = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(callerSkipOffset),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}

func (handle DailyHandle) syncer() io.Writer {
	logs, err := rotatelogs.New(
		strings.Replace(handle.path, ".log", "-%Y-%m-%d.log", 1),
		rotatelogs.WithLinkName(handle.path),
		rotatelogs.WithMaxAge(time.Duration(handle.daily)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return logs
}

func (handle DailyHandle) Debug(msg string, fields ...zap.Field) {
	handle.writer.Debug(msg, fields...)
}

func (handle DailyHandle) Debugf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Debugf(msg, fields...)
}

func (handle DailyHandle) Info(msg string, fields ...zap.Field) {
	handle.writer.Info(msg, fields...)
}

func (handle DailyHandle) Infof(msg string, fields ...interface{}) {
	handle.writer.Sugar().Infof(msg, fields...)
}

func (handle DailyHandle) Warn(msg string, fields ...zap.Field) {
	handle.writer.Warn(msg, fields...)
}

func (handle DailyHandle) Warnf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Warnf(msg, fields...)
}

func (handle DailyHandle) Error(msg string, fields ...zap.Field) {
	handle.writer.Error(msg, fields...)
}

func (handle DailyHandle) Errorf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Errorf(msg, fields...)
}

func (handle DailyHandle) DPanic(msg string, fields ...zap.Field) {
	handle.writer.DPanic(msg, fields...)
}

func (handle DailyHandle) DPanicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().DPanicf(msg, fields...)
}

func (handle DailyHandle) Panic(msg string, fields ...zap.Field) {
	handle.writer.Panic(msg, fields...)
}

func (handle DailyHandle) Panicf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Panicf(msg, fields...)
}

func (handle DailyHandle) Fatal(msg string, fields ...zap.Field) {
	handle.writer.Fatal(msg, fields...)
}

func (handle DailyHandle) Fatalf(msg string, fields ...interface{}) {
	handle.writer.Sugar().Fatalf(msg, fields...)
}
