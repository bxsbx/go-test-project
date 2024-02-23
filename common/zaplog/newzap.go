package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func InitZap(level zapcore.Level, write io.Writer) *zap.Logger {
	//配置日志输入位置
	writeSyncer := zapcore.AddSync(write)
	//配置日志打印格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig) //美化格式
	//encoder := zapcore.NewJSONEncoder(encoderConfig) //json格式
	//配置日志级别
	core := zapcore.NewCore(encoder, writeSyncer, level)

	//打印调用函数
	caller := zap.AddCaller()
	//打印调用函数时，跳过几个函数
	skip := zap.AddCallerSkip(1)
	//打印调用栈信息（error级别才打印）
	stacktrace := zap.AddStacktrace(zap.ErrorLevel)
	//获取日志对象
	return zap.New(core, caller, skip, stacktrace)
}
