package global

import (
	"StandardProject/common/zaplog"
	"go.uber.org/zap"
	"os"
)

func Init() {
	Logger = zaplog.InitZap(zap.InfoLevel,os.Stdout)
}
