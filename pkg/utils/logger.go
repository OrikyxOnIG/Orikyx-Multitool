package utils

import (
	"fmt"
	"log"
	"os"
)

func SetupLogger() *log.Logger {
	logFile, err := os.OpenFile("tool.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Could not open log file:", err)
		return nil
	}
	logger := log.New(logFile, "", log.LstdFlags)
	return logger
}

func LogMessage(message string) {
	logger := SetupLogger()
	if logger != nil {
		logger.Println(message)
	}
}
