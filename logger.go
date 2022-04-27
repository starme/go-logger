package logger

import "logger/channels"

type LogConfig struct {
	def	   string
	channels map[string]channels.ChannelConf
}

type DailyLogConf struct {
	path   string
	level  string
	daily  int
}

func (conf DailyLogConf) GetFileName() string { return conf.path }

func (conf DailyLogConf) GetLevel() string { return conf.level }


type SimpleLogConf struct {
	path  string
	level string
}

func (conf SimpleLogConf) GetFileName() string { return conf.path }

func (conf SimpleLogConf) GetLevel() string { return conf.level }
