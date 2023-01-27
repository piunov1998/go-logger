package logger

import (
	"fmt"
	"log"
	"os"
)

var Colors = map[uint32]string{
	DebugLevel:   "\u001b[37m",
	InfoLevel:    "\u001b[34m",
	WarningLevel: "\u001b[33m",
	ErrorLevel:   "\u001b[31;1m",
	FatalError:   "\u001b[4;1;38;5;160m",
}

type ColorLogger struct {
	logger *log.Logger
	level  uint32
}

func (l *ColorLogger) SetLevel(level uint32) {
	l.level = level
}

func (l *ColorLogger) log(msg string, level uint32) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("%s%s", Colors[level], prefix))
	l.logger.Print(msg + "\u001b[0m")
	l.logger.SetPrefix(prefix)
}

func (l *ColorLogger) logf(format string, level uint32, params []any) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("%s%s", Colors[level], prefix))
	l.logger.Printf(format+"\u001b[0m", params...)
	l.logger.SetPrefix(prefix)
}

func (l *ColorLogger) Debug(msg string) {
	l.log(msg, DebugLevel)
}

func (l *ColorLogger) Info(msg string) {
	l.log(msg, InfoLevel)
}

func (l *ColorLogger) Warn(msg string) {
	l.log(msg, WarningLevel)
}

func (l *ColorLogger) Error(msg string) {
	l.log(msg, ErrorLevel)
}

func (l *ColorLogger) Fatal(msg string) {
	l.log(msg, FatalError)
	os.Exit(1)
}

func (l *ColorLogger) Debugf(format string, params ...any) {
	l.logf(format, DebugLevel, params)
}

func (l *ColorLogger) Infof(format string, params ...any) {
	l.logf(format, InfoLevel, params)
}

func (l *ColorLogger) Warnf(format string, params ...any) {
	l.logf(format, WarningLevel, params)
}

func (l *ColorLogger) Errorf(format string, params ...any) {
	l.logf(format, ErrorLevel, params)
}

func (l *ColorLogger) Fatalf(format string, params ...any) {
	l.logf(format, FatalError, params)
	os.Exit(1)
}
