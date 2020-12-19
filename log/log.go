package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

type LogBody struct {
	Name    string
	Message string
	Level   string
}

var Level = INFO

func StringToNumber(s string) int {
	switch {
	case s == DEBUG:
		return 1
	case s == INFO:
		return 2
	case s == WARNING:
		return 3
	case s == ERROR:
		return 4
	default:
		return 0
	}
}

func SetLevel(level string) {
	Level = level
}

func PrintLog(logBody LogBody) {
	fmt.Printf(
		"Log | %s | %7s | %25s | %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		logBody.Level,
		logBody.Name,
		strings.ReplaceAll(logBody.Message, "\n", ""),
	)
}

func AddLog(logBody LogBody) {
	if StringToNumber(logBody.Level) >= StringToNumber(Level) {
		PrintLog(logBody)
	}
}
