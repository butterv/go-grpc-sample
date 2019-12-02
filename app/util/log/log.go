package log

import (
	"context"
	"io"
	"runtime"

	"github.com/sirupsen/logrus"
)

// CtxRequestIDKeyType is a new type for CtxRequestIDKey
type CtxRequestIDKeyType int

const (
	// CtxRequestIDKey is a key to get request id from context.
	CtxRequestIDKey CtxRequestIDKeyType = iota + 1
)

// Logger is a logger.
type Logger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
}

// New creates a new logger.
func New(ctx context.Context) *Logger {
	requestID, ok := ctx.Value(CtxRequestIDKey).(string)
	if !ok {
		requestID = "<unknown>"
	}

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		logger: l,
		entry:  l.WithField("request_id", requestID),
	}
}

// SetOutput sets outputs.
func (l *Logger) SetOutput(out io.Writer) {
	l.logger.SetOutput(out)
}

// WithPrefix sets prefix.
func (l *Logger) WithPrefix(prefix string) *Logger {
	return &Logger{
		logger: l.logger,
		entry:  l.entry.WithField("prefix", prefix),
	}
}

// WithFields sets several fields.
func (l *Logger) WithFields(m map[string]interface{}) *Logger {
	return &Logger{
		logger: l.logger,
		entry:  l.entry.WithFields(m),
	}
}

// withFileLine takes filename and line.
// This can be a bit slow (+40~60% time for logging), so
// call only when the log level is enabled.
func (l *Logger) withFileLine() *logrus.Entry {
	_, file, line, _ := runtime.Caller(2)

	return l.entry.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	})
}

// Trace emits TRACE level log.
func (l *Logger) Trace(msg string) {
	if l.logger.IsLevelEnabled(logrus.TraceLevel) {
		l.withFileLine().Trace(msg)
	}
}

// Tracef emits TRACE level log with formatter.
func (l *Logger) Tracef(format string, args ...interface{}) {
	if !l.logger.IsLevelEnabled(logrus.TraceLevel) {
		return
	}
	l.withFileLine().Tracef(format, args...)
}

// Debug emits DEBUG level log.
func (l *Logger) Debug(msg string) {
	if !l.logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}
	l.withFileLine().Debug(msg)
}

// Debugf emits DEBUG level log with formatter
func (l *Logger) Debugf(format string, args ...interface{}) {
	if !l.logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}
	l.withFileLine().Debugf(format, args...)
}

// Info emits INFO level log.
func (l *Logger) Info(msg string) {
	if !l.logger.IsLevelEnabled(logrus.InfoLevel) {
		return
	}
	l.withFileLine().Info(msg)
}

// Infof emits INFO level log with formatter.
func (l *Logger) Infof(format string, args ...interface{}) {
	if !l.logger.IsLevelEnabled(logrus.InfoLevel) {
		return
	}
	l.withFileLine().Infof(format, args...)
}

// Warn emits WARN level log.
func (l *Logger) Warn(msg string) {
	if !l.logger.IsLevelEnabled(logrus.WarnLevel) {
		return
	}
	l.withFileLine().Warn(msg)
}

// Warnf emits WARN level log.
func (l *Logger) Warnf(format string, args ...interface{}) {
	if !l.logger.IsLevelEnabled(logrus.WarnLevel) {
		return
	}
	l.withFileLine().Warnf(format, args...)
}

// Error emits ERROR level log.
func (l *Logger) Error(msg string) {
	if !l.logger.IsLevelEnabled(logrus.ErrorLevel) {
		return
	}
	l.withFileLine().Error(msg)
}

// Errorf emits ERROR level log.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if !l.logger.IsLevelEnabled(logrus.ErrorLevel) {
		return
	}
	l.withFileLine().Errorf(format, args...)
}
