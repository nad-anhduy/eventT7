package common

import (
	"os"
	"time"
)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func GetDateLog() string {
	return time.Now().UTC().Format("2006-01-02")
}
