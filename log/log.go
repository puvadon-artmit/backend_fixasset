package log

import (
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func LogInit(app *fiber.App) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	Log = logger
}

func Info(message string, fields ...zapcore.Field) {
	Log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	Log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		Log.Error(v.Error(), fields...)
	case string:
		Log.Error(v, fields...)
	}
}
