package logger

import "fmt"

type stdoutLogger struct {
}

func (l *stdoutLogger) internalPrint(args ...interface{}) {
	fmt.Println(args...)
}
func (l *stdoutLogger) internalPrintf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// Trace logs a message at level Trace on the standard logger.
func (l *stdoutLogger) Trace(args ...interface{}) {
	l.internalPrint(args...)
}

// Debug logs a message at level Debug on the standard logger.
func (l *stdoutLogger) Debug(args ...interface{}) {
	l.internalPrint(args...)
}

// Print logs a message at level Info on the standard logger.
func (l *stdoutLogger) Print(args ...interface{}) {
	l.internalPrint(args...)
}

// Info logs a message at level Info on the standard logger.
func (l *stdoutLogger) Info(args ...interface{}) {
	l.internalPrint(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l *stdoutLogger) Warn(args ...interface{}) {
	l.internalPrint(args...)
}

// Error logs a message at level Error on the standard logger.
func (l *stdoutLogger) Error(args ...interface{}) {
	l.internalPrint(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l *stdoutLogger) Panic(args ...interface{}) {
	l.internalPrint(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *stdoutLogger) Fatal(args ...interface{}) {
	l.internalPrint(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func (l *stdoutLogger) Tracef(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l *stdoutLogger) Debugf(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (l *stdoutLogger) Printf(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (l *stdoutLogger) Infof(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l *stdoutLogger) Warnf(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l *stdoutLogger) Errorf(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l *stdoutLogger) Panicf(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *stdoutLogger) Fatalf(format string, args ...interface{}) {
	l.internalPrintf(format, args...)
}
