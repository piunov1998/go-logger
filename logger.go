package logger

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

const (
	DebugLevel = 1 << iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalError
)

var levels = map[string]uint32{
	"debug":   DebugLevel,
	"info":    InfoLevel,
	"warning": WarningLevel,
	"warn":    WarningLevel,
	"error":   ErrorLevel,
	"err":     ErrorLevel,
	"fatal":   FatalError,
}

type Config struct {
	Colors   bool   `yaml:"colors"`
	LogLevel string `yaml:"log_level"`
}

var BasicConfig = Config{
	Colors:   true,
	LogLevel: "info",
}

type Logger interface {
	SetLevel(level uint32)
	Log(level uint32, msg string)
	Logf(level uint32, format string, params ...any)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Debugf(format string, params ...any)
	Infof(format string, params ...any)
	Warnf(format string, params ...any)
	Errorf(format string, params ...any)
	Fatalf(format string, params ...any)
}

func New(issuer any, config *Config) Logger {
	var name string

	switch issuer.(type) {
	default:
		name = reflect.TypeOf(issuer).Name()
	case string:
		name = issuer.(string)
	}
	flags := log.Ldate + log.Ltime + log.Lmsgprefix
	innerLogger := log.New(os.Stdout, fmt.Sprintf("[%s] ", name), flags)

	if BasicConfig.LogLevel == "" {
		BasicConfig.LogLevel = "info"
	}

	if config == nil {
		config = &BasicConfig
	}

	if config.LogLevel == "" {
		config.LogLevel = BasicConfig.LogLevel
	}

	level, ok := levels[config.LogLevel]
	if !ok {
		panic("неверный уровень логов")
	}

	if config.Colors {
		return &ColorLogger{
			logger: innerLogger,
			level:  level,
		}
	}
	return &TextLogger{
		logger: innerLogger,
		level:  level,
	}
}
