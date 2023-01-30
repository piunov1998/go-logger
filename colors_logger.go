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

func (l *ColorLogger) Log(level uint32, msg string) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("%s%s", Colors[level], prefix))
	l.logger.Print(msg + "\u001b[0m")
	l.logger.SetPrefix(prefix)
}

func (l *ColorLogger) Logf(level uint32, format string, params ...any) {
	if level < l.level {
		return
	}
	prefix := l.logger.Prefix()
	l.logger.SetPrefix(fmt.Sprintf("%s%s", Colors[level], prefix))
	l.logger.Printf(format+"\u001b[0m", params...)
	l.logger.SetPrefix(prefix)
}

func (l *ColorLogger) Debug(msg string) {
	l.Log(DebugLevel, msg)
}

func (l *ColorLogger) Info(msg string) {
	l.Log(InfoLevel, msg)
}

func (l *ColorLogger) Warn(msg string) {
	l.Log(WarningLevel, msg)
}

func (l *ColorLogger) Error(msg string) {
	l.Log(ErrorLevel, msg)
}

func (l *ColorLogger) Fatal(msg string) {
	l.Log(FatalError, msg)
	os.Exit(1)
}

func (l *ColorLogger) Debugf(format string, params ...any) {
	l.Logf(DebugLevel, format, params...)
}

func (l *ColorLogger) Infof(format string, params ...any) {
	l.Logf(InfoLevel, format, params...)
}

func (l *ColorLogger) Warnf(format string, params ...any) {
	l.Logf(WarningLevel, format, params...)
}

func (l *ColorLogger) Errorf(format string, params ...any) {
	l.Logf(ErrorLevel, format, params...)
}

func (l *ColorLogger) Fatalf(format string, params ...any) {
	l.Logf(FatalError, format, params...)
	os.Exit(1)
}
