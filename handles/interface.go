package handles

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Handler interface {
	Debug(msg string, fields ...zap.Field)
	Debugf(msg string, fields ...interface{})

	Info(msg string, fields ...zap.Field)
	Infof(msg string, fields ...interface{})

	Warn(msg string, fields ...zap.Field)
	Warnf(msg string, fields ...interface{})

	Error(msg string, fields ...zap.Field)
	Errorf(msg string, fields ...interface{})

	DPanic(msg string, fields ...zap.Field)
	DPanicf(msg string, fields ...interface{})

	Panic(msg string, fields ...zap.Field)
	Panicf(msg string, fields ...interface{})

	Fatal(msg string, fields ...zap.Field)
	Fatalf(msg string, fields ...interface{})
}

func jsonEncoder() zapcore.Encoder {
	encodeConf := zap.NewProductionEncoderConfig()
	encodeConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConf.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConf.MessageKey = "message"
	encodeConf.TimeKey = "timestamp"
	return zapcore.NewJSONEncoder(encodeConf)
}

func consoleEncoder() zapcore.Encoder {
	encodeConf := zap.NewDevelopmentEncoderConfig()
	encodeConf.EncodeTime = zapcore.TimeEncoderOfLayout("[2006-01-02 15:04:05]")
	return zapcore.NewConsoleEncoder(encodeConf)
}

func level(lvl string) zapcore.Level {
	var lv zapcore.Level

	if err := lv.UnmarshalText([]byte(lvl)); err != nil {
		panic("get log lv error: " + fmt.Sprintf("%#v", err))
	}
	return lv
}
