package logger

import (
	"fmt"
	"log"
	"os"
)

var Labels = map[uint32]string{
	DebugLevel:   "DEBUG",
	InfoLevel:    "INFO",
	WarningLevel: "WARNING",
	ErrorLevel:   "ERROR",
	FatalError:   "FATAL",
}

type TextLogger struct {
	logger *log.Logger
	level  uint32
}

func (l *TextLogger) SetLevel(level uint32) {
	l.level = level
}

func (l *TextLogger) log(msg string, level uint32) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("(%s) %s", Labels[level], prefix))
	l.logger.Print(msg)
	l.logger.SetPrefix(prefix)
}

func (l *TextLogger) logf(format string, level uint32, params []any) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("(%s) %s", Labels[level], prefix))
	l.logger.Printf(format, params...)
	l.logger.SetPrefix(prefix)
}

func (l *TextLogger) Debug(msg string) {
	l.log(msg, DebugLevel)
}

func (l *TextLogger) Info(msg string) {
	l.log(msg, InfoLevel)
}

func (l *TextLogger) Warn(msg string) {
	l.log(msg, WarningLevel)
}

func (l *TextLogger) Error(msg string) {
	l.log(msg, ErrorLevel)
}

func (l *TextLogger) Fatal(msg string) {
	l.log(msg, FatalError)
	os.Exit(1)
}

func (l *TextLogger) Debugf(format string, params ...any) {
	l.logf(format, DebugLevel, params)
}

func (l *TextLogger) Infof(format string, params ...any) {
	l.logf(format, InfoLevel, params)
}

func (l *TextLogger) Warnf(format string, params ...any) {
	l.logf(format, WarningLevel, params)
}

func (l *TextLogger) Errorf(format string, params ...any) {
	l.logf(format, ErrorLevel, params)
}

func (l *TextLogger) Fatalf(format string, params ...any) {
	l.logf(format, FatalError, params)
	os.Exit(1)
}
