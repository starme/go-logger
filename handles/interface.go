package handles

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Handler interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
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