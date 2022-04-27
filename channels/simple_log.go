package channels

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)


type SimpleLog struct {
	Config ChannelConf

	writer *zap.Logger
}

func (sl *SimpleLog) New() Channel {
	sl.init()
	return sl
}

func (sl *SimpleLog) init() {
	file, _ := os.Create(sl.Config.GetFileName())
	core := zapcore.NewCore(
		encoder(),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(file)),
		level(sl.Config.GetLevel()),
	)
	sl.writer = zap.New(core, zap.AddCaller())
}

func (sl *SimpleLog) Info(msg string, fields ...zap.Field) {
	sl.writer.Info(msg, fields...)
}