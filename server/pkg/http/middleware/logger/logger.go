package logger

import (
	"os"

	"github.com/ZyoGo/Backend-Challange/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// use singleton pattern for logger
var L *zap.Logger

func GetLogger() *zap.Logger {
	if L == nil {
		logger := InitLogger()
		defer logger.Sync()

		undo := zap.ReplaceGlobals(logger)
		defer undo()

		L = zap.L()
	}

	return L
}

func ToZapLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func InitLogger() *zap.Logger {
	consoleEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleErrors := zapcore.AddSync(os.Stdout)
	logLevel := ToZapLogLevel(config.GetConfig().App.LogLevel)

	core := zapcore.NewCore(
		consoleEncoder,
		consoleErrors,
		logLevel,
	)
	return zap.New(core)
}
