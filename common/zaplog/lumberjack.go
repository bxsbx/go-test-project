package zaplog

import (
	"github.com/natefinch/lumberjack"
)

func NewLumberjackLogger(filename string, maxSize, maxBackups, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件的最大尺寸 (以 MB 为单位)
		MaxBackups: maxBackups, // 保留的旧日志文件的最大数量
		MaxAge:     maxAge,     // 保留的旧日志文件的最大天数
		Compress:   compress,   // 是否压缩/归档旧日志文件
	}
}
