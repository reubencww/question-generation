package log

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZapLogger(development bool) *zap.Logger {
	if development {
		devConfig := zap.NewDevelopmentConfig()
		devConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		loggerInstance, err := devConfig.Build()
		if err != nil {
			panic(fmt.Sprintf("failed to create development config: %v", err))
		}

		return loggerInstance
	}

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(cfg)

	path := viper.GetString("app.log_path")
	if path == "" {
		path = "out.log"
	}

	fileLogger := newRotatingLogger(path)
	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(fileLogger), zap.DebugLevel)

	return zap.New(core, zap.WithCaller(true))
}

func newRotatingLogger(logName string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:  logName,
		MaxSize:   500,
		LocalTime: true,
		Compress:  false,
	}
}
