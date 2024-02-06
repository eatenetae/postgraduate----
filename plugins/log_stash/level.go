package log_stash

import "encoding/json"

type Level int

const (
	DEBUG Level = 1
	INFO  Level = 2
	WARN  Level = 3
	ERROR Level = 4
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Level) String() string {
	var str string
	switch s {
	case DEBUG:
		str = "debug"
	case INFO:
		str = "info"
	case WARN:
		str = "warn"
	case ERROR:
		str = "error"
	default:
		str = "其他"
	}
	return str
}
