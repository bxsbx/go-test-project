package middleware

import (
	"StandardProject/common/util"
	"StandardProject/common/zaplog"
	"StandardProject/gin/global"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	reqId = "req_id"
)

// 做一些公共处理，比如上下文，记录一些日志信息。
func Common() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ctx := context.Background()
		ctx = context.WithValue(ctx, reqId, util.Getuid())
		global.SetContext(c, ctx)
		c.Next()
		latency := time.Since(start) //记录耗时
		zaplog.Request(ctx, c.Request, latency.Seconds())
	}
}
