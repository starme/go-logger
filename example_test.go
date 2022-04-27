package logger

import (
	"logger/channels"
	"testing"
)

func TestNewLog(t *testing.T) {
	conf := LogConfig{
		def: "daily",
		channels: map[string]channels.ChannelConf{
			"simple": SimpleLogConf{
				path: "./log/test.log",
				level: "Info",
			},
		},
	}

	NewLog(conf)

	With("simple").Info("bbbbbbbbbb")
	With("simple").Info("cccccccccc")
	With("simple").Info("dddddddddd")
}