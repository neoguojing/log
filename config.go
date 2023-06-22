package log

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogFormat string

const (
	JSON    LogFormat = "json"
	CONSOLE LogFormat = "console"
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

	basePath             = os.Getenv("LOG_PATH")
	logFileName          = os.Args[0]
	DefaultLogFileConfig = &lumberjack.Logger{
		Filename:   filepath.Join(basePath, logFileName),
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     28, // days
		LocalTime:  true,
	}
)

type Option func(*LoggerConfig)

func WithComstomLogConfig(fileName string, MaxSize int, MaxAge int) Option {
	return func(l *LoggerConfig) {
		l.fileConfig = &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    MaxSize, // megabytes
			MaxBackups: 5,
			MaxAge:     MaxAge, // days
			LocalTime:  true,
		}
	}
}

type LoggerConfig struct {
	level      LogLevel
	fileConfig *lumberjack.Logger
	coreLevel  zap.AtomicLevel
	output     zapcore.WriteSyncer
	encoder    zapcore.Encoder
	core       zapcore.Core
}

func NewConfig() *LoggerConfig {
	l := &LoggerConfig{level: DEBUG}
	l.fileConfig = DefaultLogFileConfig
	l.coreLevel = zap.NewAtomicLevel()
	l.coreLevel.SetLevel(zap.DebugLevel)
	l.encoder = encoderConsole
	l.output = zapcore.Lock(os.Stdout)
	l.core = zapcore.NewCore(l.encoder, l.output, l.coreLevel)

	return l
}

func (l *LoggerConfig) Level(level LogLevel) *LoggerConfig {
	l.level = level
	return l
}

func (l *LoggerConfig) Rotate(opts ...Option) *LoggerConfig {
	l.fileConfig = DefaultLogFileConfig
	for _, opt := range opts {
		opt(l)
	}

	l.output = zapcore.NewMultiWriteSyncer(zapcore.AddSync(l.fileConfig), zapcore.AddSync(os.Stdout))
	l.core = zapcore.NewCore(l.encoder, l.output, l.coreLevel)
	return l
}

func (l *LoggerConfig) Format(format LogFormat) *LoggerConfig {

	switch format {
	case JSON:
		l.encoder = encoderJson

	case CONSOLE:
		l.encoder = encoderConsole
	default:
		l.encoder = encoderConsole
	}
	l.core = zapcore.NewCore(l.encoder, l.output, l.coreLevel)
	return l
}

func (l *LoggerConfig) Build() *Logger {
	log := &Logger{level: l.level}
	log.logger = zap.New(l.core, zap.AddCaller(), zap.AddStacktrace(zap.FatalLevel))
	log.logger.WithOptions(zap.AddCallerSkip(2))
	return log

}
