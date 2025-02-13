package logger

import (
	"fmt"
	"log"
	"os"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var (
	levelNames   = []string{"DEBUG", "INFO", "WARN", "ERROR"}
	currentLevel = INFO
	logger       = log.New(os.Stdout, "", 0)
)

func SetLevel(level Level) {
	if level >= DEBUG && level <= ERROR {
		currentLevel = level
	}
}

func logMessage(level Level, message string) {
	if level >= currentLevel {
		logger.Printf("[%s] %s", levelNames[level], message)
	}
}

// Debug logs the message with a debug level
func Debug(message string) {
	logMessage(DEBUG, message)
}

// Info logs the message with an info level
func Info(message string) {
	logMessage(INFO, message)
}

// Warn logs the message with a warning level
func Warn(message string) {
	logMessage(WARN, message)
}

// Error logs the message with an error level
// and exits the program with status code 1
func Error(message string) {
	logMessage(ERROR, message)
	os.Exit(1)
}

// Debugf logs the formatted message with a debug level
func Debugf(format string, args ...interface{}) {
	logMessage(DEBUG, fmt.Sprintf(format, args...))
}

// Infof logs the formatted message with an info level
func Infof(format string, args ...interface{}) {
	logMessage(INFO, fmt.Sprintf(format, args...))
}

// Warnf logs the formatted message with a warning level
func Warnf(format string, args ...interface{}) {
	logMessage(WARN, fmt.Sprintf(format, args...))
}

// Errorf logs the formatted message with an error level
// and exits the program with status code 1
func Errorf(format string, args ...interface{}) {
	logMessage(ERROR, fmt.Sprintf(format, args...))
	os.Exit(1)
}
