package zaplog

import "go.uber.org/zap"

func InitZap() (err error) {
	Logger, err = writeToConsole()
	return
}

// 直接输出到控制台
func writeToConsole() (*zap.Logger, error) {
	return zap.NewProductionConfig().Build(zap.AddCallerSkip(1))
}
