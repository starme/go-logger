package logger

type LogType string

const (
	Simple LogType = "simple"
	Daily  LogType = "daily"
	Std    LogType = "std"
)

type Config struct {
	Default string

	Channels []Channel
}

type Channel struct {
	Type  LogType
	Name  string
	Path  string
	Level string
	Days  int
}
