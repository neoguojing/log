package log

var (
	logger = NewLogger()
)

func SetLogger(l *Logger) {
	logger = l
}

func Debug(msg ...string) {
	logger.Debug(msg...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(msg ...string) {
	logger.Info(msg...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warning(msg ...string) {
	logger.Warning(msg...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Error(msg ...string) {
	logger.Error(msg...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
