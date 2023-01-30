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

func (l *TextLogger) Log(level uint32, msg string) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("(%s) %s", Labels[level], prefix))
	l.logger.Print(msg)
	l.logger.SetPrefix(prefix)
}

func (l *TextLogger) Logf(level uint32, format string, params ...any) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("(%s) %s", Labels[level], prefix))
	l.logger.Printf(format, params...)
	l.logger.SetPrefix(prefix)
}

func (l *TextLogger) Debug(msg string) {
	l.Log(DebugLevel, msg)
}

func (l *TextLogger) Info(msg string) {
	l.Log(InfoLevel, msg)
}

func (l *TextLogger) Warn(msg string) {
	l.Log(WarningLevel, msg)
}

func (l *TextLogger) Error(msg string) {
	l.Log(ErrorLevel, msg)
}

func (l *TextLogger) Fatal(msg string) {
	l.Log(FatalError, msg)
	os.Exit(1)
}

func (l *TextLogger) Debugf(format string, params ...any) {
	l.Logf(DebugLevel, format, params...)
}

func (l *TextLogger) Infof(format string, params ...any) {
	l.Logf(InfoLevel, format, params...)
}

func (l *TextLogger) Warnf(format string, params ...any) {
	l.Logf(WarningLevel, format, params...)
}

func (l *TextLogger) Errorf(format string, params ...any) {
	l.Logf(ErrorLevel, format, params...)
}

func (l *TextLogger) Fatalf(format string, params ...any) {
	l.Logf(FatalError, format, params...)
	os.Exit(1)
}
