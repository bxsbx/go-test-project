package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Pring(logger *zap.Logger) {
	field := zap.Any("header", "req.Header")
	// 记录不同级别的日志
	logger.Debug("This is a debug log", field)
	logger.Info("This is an info log", field)
	logger.Warn("This is a warning log", field)
	logger.Error("This is an error log", field)
}

func Pring2(logger *zap.SugaredLogger) {
	field := zap.Any("header", "req.Header")
	// 记录不同级别的日志
	logger.Debug("This is a debug log", field)
	logger.Info("This is an info log", field)
	logger.Warn("This is a warning log", field)
	logger.Error("This is an error log", field)
}

func main() {

	// 配置日志级别
	//atomicLevel := zap.NewAtomicLevel()
	//
	//// 设置日志级别为Debug
	//atomicLevel.SetLevel(zapcore.DebugLevel)
	//
	//// 创建配置
	//config := zap.NewProductionConfig()
	//config.Level = atomicLevel
	//
	//// 构建Logger
	//logger, _ := config.Build()
	//Pring(logger)

	//file, _ := os.OpenFile("filePath.text", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	//logger := log.New(os.Stdout, "prelog", log.LstdFlags)
	//logger.Println("cwqwfqwf")
	//
	//log.Println()

	writeSyncer := zapcore.AddSync(os.Stdout)

	//encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.InfoLevel))
	sugaredLogger := logger.Sugar()
	Pring(logger)
	Pring2(sugaredLogger)

}
