package logger

var logCenter = make(map[string]*Driver)


func With(name string) *Driver {
	if ch, ok := logCenter[name]; ok {
		return ch
	}
	return nil
}

func NewLog(config Config) {
	for _, config := range config.Channels {
		channel := makeChannel(config)
		if channel == nil {
			continue
		}
		logCenter[config.Name] = channel
	}
}

func makeChannel(config Channel) *Driver {
	if handle, ok := MakeHandle(config); ok {
		return &Driver{
			handler: handle,
		}
	}
	return nil
}
