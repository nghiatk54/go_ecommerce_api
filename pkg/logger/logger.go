package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/nghiatk54/goEcommerceApi/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	sync := getWriterSync(config)

	core := zapcore.NewCore(encoder, 
		sync, 
		level)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return &LoggerZap{logger}
}

// format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// write log
func getWriterSync(config setting.LoggerSetting) zapcore.WriteSyncer {
	hook := lumberjack.Logger{
		Filename: config.Filename,
		MaxSize: config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge: config.MaxAge,
		Compress: config.Compress,
	}
	syncFile := zapcore.AddSync(&hook)
	syncConsole := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}