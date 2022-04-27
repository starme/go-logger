package channels

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ChannelConf interface {
	GetFileName() string
	GetLevel()	  string
}

type Channel interface {
	Info(msg string, fields ...zap.Field)
	New() Channel
}

func encoder() zapcore.Encoder {
	encodeConf := zap.NewProductionEncoderConfig()
	encodeConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConf.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConf.MessageKey = "message"
	encodeConf.TimeKey = "timestamp"
	return zapcore.NewJSONEncoder(encodeConf)
}

func level(lvl string) zapcore.Level {
	var lv zapcore.Level

	if err := lv.UnmarshalText([]byte(lvl)); err != nil {
		panic("get log lv error: " + fmt.Sprintf("%#v", err))
	}
	return lv
}