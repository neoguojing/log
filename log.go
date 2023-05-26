package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

var (
	encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoderConsole = zapcore.NewConsoleEncoder(encoderConfig)
	encoderJson    = zapcore.NewJSONEncoder(encoderConfig)
)

type Option func(*Logger)

type Logger struct {
	level     LogLevel
	logger    *zap.Logger
	output    zapcore.WriteSyncer
	encoder   zapcore.Encoder
	core      zapcore.Core
	coreLevel zap.AtomicLevel
}

func NewLogger() *Logger {
	l := &Logger{level: DEBUG}
	l.coreLevel = zap.NewAtomicLevel()
	l.coreLevel.SetLevel(zap.DebugLevel)
	// 设置日志输出格式

	// 设置日志输出
	l.output = zapcore.Lock(os.Stdout)
	l.core = zapcore.NewCore(encoderConsole, l.output, l.coreLevel)

	// 创建 Logger
	l.logger = zap.New(l.core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return l
}

func (l *Logger) Level(level LogLevel) *Logger {

	l.level = level
	return l

}

func (l *Logger) Rotate(path string) *Logger {
	ljLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     28, // days
		LocalTime:  true,
	}
	l.output = zapcore.AddSync(ljLogger)
	l.core = zapcore.NewCore(l.encoder, l.output, l.coreLevel)
	l.logger = zap.New(l.core)
	return l
}

func (l *Logger) Format(format string) *Logger {

	switch format {
	case "json":
		l.core = zapcore.NewCore(encoderJson, l.output, zap.NewAtomicLevelAt(zap.DebugLevel))
	case "console":
		l.core = zapcore.NewCore(encoderConsole, l.output, zap.NewAtomicLevelAt(zap.DebugLevel))
	default:
		l.core = zapcore.NewCore(encoderConsole, l.output, zap.NewAtomicLevelAt(zap.DebugLevel))
	}
	l.logger = zap.New(l.core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return l
}

func (l *Logger) Debug(msg string) {
	if l.level <= DEBUG {
		l.logger.Debug(msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.level <= INFO {
		l.logger.Info(msg)
	}
}

func (l *Logger) Warning(msg string) {
	if l.level <= WARNING {
		l.logger.Warn(msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.level <= ERROR {
		l.logger.Error(msg)
	}
}
