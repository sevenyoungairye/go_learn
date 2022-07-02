package yaml

import (
	"gopkg.in/yaml.v3"
)

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
)

var LogLevelMap = map[string]int{}

func init() {
	LogLevelMap[Debug] = 1
	LogLevelMap[Info] = 2
	LogLevelMap[Warn] = 3
	LogLevelMap[Error] = 4
}

// L define a log config.
type L struct {
	Log struct {
		Level string `yaml:"level" json:"level"`
		// log level.
		val int
	} `yaml:"log" json:"log"`
}

func GetLogConfig() *L {
	l := L{}
	_ = yaml.Unmarshal(configData, &l)
	if "" == l.Log.Level {
		l.Log.Level = Info
	}
	l.Log.val = LogLevelMap[l.Log.Level]
	return &l
}

func (l *L) GetLogVal() int {
	return l.Log.val
}
