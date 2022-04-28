package logger

import (
	"github.com/starme/logger/handles"
)

var logCenter = make(map[string]handles.Handler)

var defDriver = ""

func With(name string) handles.Handler {
	if ch, ok := logCenter[name]; ok {
		return ch
	}

	panic("log driver is not found.")
}

func NewLog(config Config) {
	defDriver = config.Default
	for _, config := range config.Channels {
		if handle, ok := MakeHandle(config); ok {
			logCenter[config.Name] = handle
		}
	}
}
