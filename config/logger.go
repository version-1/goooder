package config

import (
	"fmt"
	"log"
)

type logLevel int

const (
	LogLevelInfo logLevel = iota
	LogLevelWarn
	LogLevelError
)

type DefaultLogger struct {
	logLevel
}

func (l *DefaultLogger) SetLogLevel(level string) {
	switch level {
	case "info":
		l.logLevel = LogLevelInfo
	case "warn":
		l.logLevel = LogLevelWarn
	case "error":
		l.logLevel = LogLevelError
	default:
		l.logLevel = LogLevelInfo
	}
}

func (l DefaultLogger) Infof(format string, args ...any) {
	if l.logLevel > LogLevelInfo {
		return
	}
	log.Printf(fmt.Sprintf("[INFO] %s", format), args...)
}

func (l DefaultLogger) Warnf(format string, args ...any) {
	if l.logLevel > LogLevelWarn {
		return
	}
	log.Printf(fmt.Sprintf("[WARN] %s", format), args...)
}

func (l DefaultLogger) Errorf(format string, args ...any) {
	if l.logLevel > LogLevelError {
		return
	}
	log.Printf(fmt.Sprintf("[ERROR] %s", format), args...)
}

func (l DefaultLogger) Fatalf(format string, args ...any) {
	log.Printf(fmt.Sprintf("[FATAL] %s", format), args...)
	err := fmt.Errorf(format, args...)
	panic(err)
}
