package global

import (
	"StandardProject/gin/config"
	"go.uber.org/zap"
)

var (
	AllConfigInfo config.AllConfig
	Logger *zap.Logger
)
