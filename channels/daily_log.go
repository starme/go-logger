package channels

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type DailyLog struct {
	Config ChannelConf

	writer *zap.Logger
}

func (dl *DailyLog) New() Channel {
	dl.init()
	return dl
}

func (dl *DailyLog) init() {
	file, _ := os.Create(dl.Config.GetFileName())
	core := zapcore.NewCore(
		encoder(),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(file)),
		level(dl.Config.GetLevel()),
	)
	dl.writer = zap.New(core, zap.AddCaller())
}

func (dl *DailyLog) Info(msg string, fields ...zap.Field) {
	dl.writer.Info(msg, fields...)
}