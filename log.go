package logger

import (
	"github.com/starme/logger/channels"
)

var channelCenter map[string]channels.Channel

func With(name string) channels.Channel {
	if ch, ok := channelCenter[name]; ok {
		return ch
	}
	return nil
}

func NewLog(conf LogConfig) {
	channelCenter = make(map[string]channels.Channel)
	for name, config := range conf.channels {
		channel := makeChannel(name, config)
		if channel == nil {
			continue
		}
		channelCenter[name] = channel
	}
}

func makeChannel(name string, config channels.ChannelConf) channels.Channel {
	var channel channels.Channel
	switch name {
	case "daily":
		channel = &channels.DailyLog{Config: config}
	case "simple":
		channel = &channels.SimpleLog{Config: config}
	default:
		channel = nil
	}
	if channel == nil {
		return nil
	}
	return channel.New()
}
