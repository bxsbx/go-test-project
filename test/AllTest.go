package main

import (
	"StandardProject/common/zaplog"
	"go.uber.org/zap"
	"sync"
	"unicode/utf8"

	"github.com/natefinch/lumberjack"
)

// 脱敏姓名
func desensitizeName(name string) string {
	runeCount := utf8.RuneCountInString(name)
	if runeCount == 0 {
		return ""
	}
	first, _ := utf8.DecodeRuneInString(name)
	desensitizedName := string(first)
	for i := 1; i < runeCount-1; i++ {
		desensitizedName += "*"
	}
	if runeCount > 2 {
		last, _ := utf8.DecodeLastRuneInString(name)
		desensitizedName += string(last)
	} else {
		desensitizedName += "*"
	}
	return desensitizedName
}

type Al struct {
	sync.RWMutex
	A int
}

// NVUZCk1HwYEovkhSQhDwonCFsIsGgRom
type rw[T int8 | int16] interface {
	writer(t T)
	reader() T
}

type Number struct {
	buf int8
}

func (n *Number) writer(b int8) {
	n.buf = b
}

func (n *Number) reader() int8 {
	return n.buf
}

func main() {
	logger := &lumberjack.Logger{
		Filename:   "./logs/app.log", // 日志文件路径
		MaxSize:    1,                // 每个日志文件的最大尺寸 (以 MB 为单位)
		MaxBackups: 3,                // 保留的旧日志文件的最大数量
		MaxAge:     28,               // 保留的旧日志文件的最大天数
		Compress:   true,             // 是否压缩/归档旧日志文件
	}
	zapLogger := zaplog.InitZap(zap.InfoLevel, logger)

	for i := 0; i < 10000; i++ {
		zapLogger.Info("csac", zap.String("1", "vqwwcsjiaaadwqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqaaaaaaaaaaaaaaaaaaiwevjwiehvwiehvowehvowhoeuvhowhevowheovhweooooooooooooooooooooooooooooooooooooowehooooooooovhuffffffffffdhu"))
	}
}
