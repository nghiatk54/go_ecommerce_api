package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Sugar
	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello, name: %s, age: %d", "nghiatk54", 28)
	// Logger
	logger := zap.NewExample()
	logger.Info("Hello new example", zap.String("name", "nghiatk54"), zap.Int("age", 28))
	// Development
	logger, _ = zap.NewDevelopment()
	logger.Info("Hello new development")
	// Production
	logger, _ = zap.NewProduction()
	logger.Info("Hello new production")
	// Custom log
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger = zap.New(core, zap.AddCaller())
	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// Format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// Set writer to file log
func getWriterSync() zapcore.WriteSyncer {
	err := os.MkdirAll("./log", os.ModePerm)
	if err != nil {
		log.Fatalf("Error create dir: %v", err)
	}
	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("Error open file: %v", err)
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
