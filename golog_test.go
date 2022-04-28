package logger

import (
	"testing"
)

func TestNewLog(t *testing.T) {
	conf := Config{
		Default: "test",
		Channels: []Channel{
			{
				Type: "simple",
				Name: "test",
				Path: "./log/test.log",
				Level: "Info",
			},
			{
				Type: "daily",
				Name: "my_log",
				Path: "./log/test_02.log",
				Level: "Info",
				Days: 7,
			},
		},
	}

	NewLog(conf)

	With("my_log").Info("bbbbbbbbbb")
	With("my_log").Info("cccccccccc")
	With("my_log").Info("dddddddddd")

	With("test").Info("bbbbbbbbbb")
	With("test").Info("cccccccccc")
	With("test").Info("dddddddddd")

}