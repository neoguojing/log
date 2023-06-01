package log

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	level  LogLevel
	logger *zap.Logger
}

func NewLogger() *Logger {
	config := NewConfig().Rotate()
	return config.Build()
}

func (l *Logger) Debug(msg ...string) {
	l.log(DEBUG, msg...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(DEBUG, format, args...)
}

func (l *Logger) Info(msg ...string) {
	l.log(INFO, msg...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func (l *Logger) Warning(msg ...string) {
	l.log(WARNING, msg...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logf(WARNING, format, args...)
}

func (l *Logger) Error(msg ...string) {
	l.log(ERROR, msg...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(ERROR, format, args...)
}

	
func (l *Logger) Fatal(msg ...string) {
	l.log(FATAL, msg...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logf(FATAL, format, args...)
}



func (l *Logger) log(level LogLevel, msg ...string) {
	if l.level <= level {
		switch level {
		case DEBUG:
			l.logger.Debug(strings.Join(msg, " "))
		case INFO:
			l.logger.Info(strings.Join(msg, " "))
		case WARNING:
			l.logger.Warn(strings.Join(msg, " "))
		case ERROR:
			l.logger.Error(strings.Join(msg, " "))
		case FATAL:
			l.logger.Fatal(strings.Join(msg, " "))
		}
	}
}

func (l *Logger) logf(level LogLevel, format string, args ...interface{}) {
	if l.level <= level {
		switch level {
		case DEBUG:
			l.logger.Debug(fmt.Sprintf(format, args...))
		case INFO:
			l.logger.Info(fmt.Sprintf(format, args...))
		case WARNING:
			l.logger.Warn(fmt.Sprintf(format, args...))
		case ERROR:
			l.logger.Error(fmt.Sprintf(format, args...))
		case FATAL:
			l.logger.Fatal(fmt.Sprintf(format, args...))
		}
	}
}


