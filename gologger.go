package logger

type Config struct {
	Default string

	Channels []Channel
}

type Channel struct {
	Type	string
	Name	string
	Path	string
	Level	string
	Days	int
}