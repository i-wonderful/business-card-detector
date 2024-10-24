package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Field struct {
	Key   string
	Value interface{}
}

type Logger struct {
	logger *zap.Logger
}

func NewLogger(name string, level string, isProduction bool) (*Logger, error) {
	lev, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %v", err)
	}

	var logConfig zap.Config
	if isProduction {
		logConfig = zap.NewProductionConfig()
	} else {
		logConfig = zap.NewDevelopmentConfig()
	}

	logConfig.Level = lev
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	logConfig.DisableCaller = true

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logConfig.EncoderConfig = encoderCfg

	logger, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{logger: logger.Named(name)}, nil
}

func (l *Logger) Debug(message string, fields ...Field) {
	zapFields := l.toZap(fields)
	l.logger.Debug(message, zapFields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	zapFields := l.toZap(fields)
	l.logger.Info(msg, zapFields...)
}

func (l *Logger) Info1(msg string, key string, value interface{}) {
	l.logger.Info(msg, zap.Any(key, value))
}

func (l *Logger) Info2(msg string, key1 string, value1 interface{}, key2 string, value2 interface{}) {
	l.logger.Info(msg, zap.Any(key1, value1), zap.Any(key2, value2))
}

func (l *Logger) Error(msg string, err error, fields ...Field) {
	zapFields := l.toZap(fields)
	zapFields = append(zapFields, zap.Error(err))
	l.logger.Error(msg, zapFields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	zapFields := l.toZap(fields)
	l.logger.Warn(msg, zapFields...)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal(msg, zap.Error(err))
}

func (l *Logger) toZap(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

func createLogger() *zap.SugaredLogger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console", //"json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return zap.Must(config.Build()).Sugar()
}
