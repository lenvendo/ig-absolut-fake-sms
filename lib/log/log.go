package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLoggerFromConfig(conf Config, isDebug bool) (*zap.Logger, error) {
	var zapConfig zap.Config
	if isDebug {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapConfig.Level = getLevel(conf.Level)
	zapConfig.Encoding = conf.Encoding
	zapConfig.OutputPaths = []string{"stdout"}

	if len(conf.Outputs) != 0 {
		zapConfig.OutputPaths = conf.Outputs
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

type Config struct {
	Level    string
	Encoding string
	Outputs  []string
}

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

func getLevel(level string) zap.AtomicLevel {
	switch level {
	case DebugLevel:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case ErrorLevel:
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case InfoLevel:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case WarnLevel:
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	default:
		panic("unknown log level")
	}
}
