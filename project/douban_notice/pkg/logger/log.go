package logger

import (
	"bytes"
	"github.com/fatih/color"
	"log"
	"top.lel.dn/main/pkg/yaml"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "lel-dn-logger: ", log.LstdFlags)
)

var config *yaml.L

func init() {
	config = yaml.GetLogConfig()
	log.Printf("log config init... using level: %s", config.Log.Level)
}

func Debug(msg string) {
	// 根据日志等级开启debug.
	if yaml.LogLevelMap[yaml.Debug] >= config.GetLogVal() {
		buf.Reset()
		logger.Printf("debug: %s", msg)
		c := color.New(color.FgHiCyan)
		_, err := c.Print(&buf)
		if err != nil {
			return
		}
	}
}

// Info https://pkg.go.dev/github.com/fatih/color
func Info(msg string) {
	if yaml.LogLevelMap[yaml.Info] >= config.GetLogVal() {
		buf.Reset()
		logger.Printf("info: %s", msg)

		c := color.New(color.FgHiGreen)
		_, err := c.Print(&buf)
		if err != nil {
			return
		}
	}
}

func Warn(msg string) {
	if yaml.LogLevelMap[yaml.Warn] >= config.GetLogVal() {
		buf.Reset()
		logger.Printf("warn: %s", msg)

		c := color.New(color.FgYellow)
		_, err := c.Print(&buf)
		if err != nil {
			return
		}
	}
}

func Error(msg string) {
	if yaml.LogLevelMap[yaml.Error] >= config.GetLogVal() {
		buf.Reset()
		logger.SetFlags(log.Lshortfile)
		logger.Fatalf("error: %s", msg)
	}
}
