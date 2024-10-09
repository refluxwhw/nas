package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is the interface for Logger types
type Logger interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	Print(args ...interface{})

	Panicf(fmt string, args ...interface{})
	Fatalf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Infof(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
	Tracef(fmt string, args ...interface{})
	Printf(fmt string, args ...interface{})
}

// standard logger
var std Logger = &stdoutLogger{}

func InitLogrus() {
	// TODO: add configure

	logFile := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // 是否压缩/归档旧文件
	}

	log := logrus.New()
	log.SetOutput(logFile)
	log.SetLevel(logrus.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerFirst:     true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})
	std = log
}

// SetLogger set standard logger
func SetLogger(log Logger) {
	std = log
}

// GetLogger get standard logger
func GetLogger() Logger {
	return std
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	std.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	std.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}
