package handles

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"strings"
	"time"
)

type DailyHandle struct {
	path	string
	level	string
	daily	int

	writer	*zap.Logger
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
		encoder(),
		zapcore.AddSync(handle.syncer()),
		level(handle.level),
	)
	handle.writer = zap.New(core, zap.AddCaller())
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

func (handle DailyHandle) Info(msg string, fields ...zap.Field) {
	handle.writer.Info(msg, fields...)
}
