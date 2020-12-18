package logger

import (
	"errors"
)

// A global variable so that log functions can be directly accessed
var log Logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	//InstanceZapLogger is an instance of the zap logger
	InstanceZapLogger int = iota
	//InstanceLogrusLogger is an instance of the logrus logger
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool   `json:"enableConsole"`
	ConsoleLevel      string `json:"consoleLevel"`
	ConsoleJSONFormat bool   `json:"consoleJSONformat"`
	EnableFile        bool   `json:"enableFile"`
	FileLevel         string `json:"fileLevel"`
	FileJSONFormat    bool   `json:"fileJSONformat"`
	FileLocation      string `json:"fileLocation"`
	MaxSize           int    `json:"maxSize"`
	Compress          bool   `json:"compress"`
	MaxAge            int    `json:"maxAge"`
}

//NewLogger returns an instance of logger
func NewLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil

	case InstanceLogrusLogger:
		logger, err := newLogrusLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil

	default:
		return errInvalidLoggerInstance
	}
}

//Debugf is the debug level logger instance
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

//Infof is the info level logger instance
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

//Warnf is the warn level logger instance
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

//Errorf is the error level logger instance
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

//Fatalf is the fatal level logger instance
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

//Panicf is the panic level logger instance
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

//WithFields allows the creation of custom fields as key value pairs.
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
