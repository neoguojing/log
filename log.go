package log

import (
	"go.uber.org/zap"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type Logger struct {
	level  LogLevel
	logger *zap.Logger
}

func NewLogger() *Logger {
	config := NewConfig().Rotate()
	return config.Build()
}

func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Warning(msg string) {
	l.log(WARNING, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

func (l *Logger) log(level LogLevel, msg string) {
	if l.level <= level {
		switch level {
		case DEBUG:
			l.logger.Debug(msg)
		case INFO:
			l.logger.Info(msg)
		case WARNING:
			l.logger.Warn(msg)
		case ERROR:
			l.logger.Error(msg)
		}
	}
}
