package logger

import (
	// "os"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger(config Configuration) (Logger, error) {
	cores := []zapcore.Core{}

	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		// writer := zapcore.Lock(os.Stdout)
		writer := zapcore.AddSync(colorable.NewColorableStdout())
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if config.EnableFile {
		level := getZapLevel(config.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileLocation,
			MaxSize:  config.MaxSize, //megabytes
			Compress: config.Compress,
			MaxAge:   config.MaxAge, //days
		})
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}
//Debugf is the debug level logger call
func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

//Infof is the info level logger call
func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}
//Warnf is the warning level logger call
func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

//Errorf is the warning error logger call
func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}
//Fatalf is the fatal level logger call
func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

//Panicf is the panic level logger call
func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Panicf(format, args...)
}

//WithFields assigns needed fields to the logger
func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
