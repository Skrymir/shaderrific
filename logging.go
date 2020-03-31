package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func configureLogging() *zap.SugaredLogger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./shaderiffic.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     15, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(logEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	return zap.New(core).Sugar()
}

func logEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
