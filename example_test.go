package logger

import (
	"go.uber.org/zap"
)

func ExampleNewLog() {
	conf := Config{
		Default: "Example",
		Channels: []Channel{
			{
				Type:         "simple",
				Name:         "Example",
				Path:         "./log/Example.log",
				Level:        "Info",
				CallerEnable: true,
			},
			{
				Type:         "daily",
				Name:         "my_log",
				Path:         "./log/Example_02.log",
				Level:        "Info",
				MaxAge:       7,
				CallerEnable: false,
			},
		},
	}

	NewLog(conf)
}

func ExampleWith() {
	ExampleNewLog()
	With("simpleLog")
}

func ExampleDebug() {
	ExampleNewLog()
	Debug("this is debug log.", zap.String("context", "lalala..."))
}

func ExampleInfo() {
	ExampleNewLog()
	Info("this is info log.")
}

func ExampleWarn() {
	ExampleNewLog()
	Warn("this is warning log.")
}

func ExampleError() {
	ExampleNewLog()
	Error("this is error log.")
}

func ExampleDPanic() {
	ExampleNewLog()
	DPanic("this is panic log.")
}

func ExamplePanic() {
	ExampleNewLog()
	Panic("this is panic log.")
}

func ExampleFatal() {
	ExampleNewLog()
	Fatal("this is fatal log.")
}

func ExampleDebugf() {
	ExampleNewLog()
	Debugf("%s this is debug log.", "lalala...")
}

func ExampleInfof() {
	ExampleNewLog()
	Infof("this is info log.")
}

func ExampleWarnf() {
	ExampleNewLog()
	Warnf("this is warning log.")
}
