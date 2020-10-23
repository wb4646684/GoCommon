package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

type LogBody struct {
	Name    string
	Message string
	Level   string
}

func PrintLog(logBody LogBody) {
	fmt.Printf(
		"MDOpsGW | %s | %7s | %25s | %s \n",
		time.Now().Format("2006-01-02 15:04:05"),
		logBody.Level,
		logBody.Name,
		strings.ReplaceAll(logBody.Message, "\n", ""),
	)
}
