package core

import (
	"ToriBackend/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitZapLogger(env string) *zap.Logger {
	if env == "release" {
		return initZapFile()
	} else {
		return initZapConsole()
	}
}

func initZapFile() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   global.Config.Zap.LogFilename, // save log file in the project root directory
		MaxSize:    300,                           // MB
		MaxBackups: 3,                             // max file num to keep
		MaxAge:     31,                            // all file older than max age would be deleted, despite max backups
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)
	// Make the global logger able to write into json, must invoke
	zap.ReplaceGlobals(logger)
	return logger
}

func initZapConsole() *zap.Logger {
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(getEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	return logger
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "function",
		StacktraceKey:  global.Config.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
