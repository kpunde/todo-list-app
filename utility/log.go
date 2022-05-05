package utility

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime/debug"
	"time"
)

var cfg = zap.Config{
	Encoding:         "json",
	Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	},
}
var logger, _ = cfg.Build()

func GetLogger() *zap.Logger {
	return logger
}

func Log(level zapcore.Level, msg interface{}) {
	switch level {
	case zap.ErrorLevel:
		logger.Error("Error Found",
			zap.Time("time", time.Now()),
			zap.Any("error", msg),
			zap.String("stack", string(debug.Stack())))
	case zap.DebugLevel:
		logger.Debug("Error Found",
			zap.Time("time", time.Now()),
			zap.Any("message", msg))

	}
}
