package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewLog(t *testing.T) {
	conf := Config{
		Default: "stdLog",
		Channels: []Channel{
			{
				Type: "std",
				Name: "stdLog",
			},
			{
				Type:  "simple",
				Name:  "simpleLog",
				Level: "Info",
			},
			{
				Type:  "daily",
				Name:  "dailyLog",
				Level: "Info",
			},
		},
	}
	NewLog(conf)
}

func TestWith(t *testing.T) {
	TestNewLog(t)
	With("simpleLog")
}

func TestDebug(t *testing.T) {
	TestNewLog(t)
	Debug("this is debug log.", zap.String("context", "lalala..."))
}

func TestInfo(t *testing.T) {
	TestNewLog(t)
	Info("this is info log.")
}

func TestWarn(t *testing.T) {
	TestNewLog(t)
	Warn("this is warning log.")
}

func TestError(t *testing.T) {
	TestNewLog(t)
	Error("this is error log.")
}

func TestDPanic(t *testing.T) {
	TestNewLog(t)
	DPanic("this is panic log.")
}

func TestPanic(t *testing.T) {
	TestNewLog(t)
	Panic("this is panic log.")
}

func TestFatal(t *testing.T) {
	TestNewLog(t)
	Fatal("this is fatal log.")
}

func TestDebugf(t *testing.T) {
	TestNewLog(t)
	Debugf("%s this is debug log.", "lalala...")
}

func TestInfof(t *testing.T) {
	TestNewLog(t)
	Infof("this is info log.")
}

func TestWarnf(t *testing.T) {
	TestNewLog(t)
	Warnf("this is warning log.")
}
