package zaplog

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

var Logger *zap.Logger

func Request(ctx context.Context, req *http.Request, latency float64) {
	fields := []zap.Field{
		zap.Any("req_id", ctx.Value("req_id")),
		zap.String("url", req.URL.String()),
		zap.String("method", req.Method),
		zap.Any("header", req.Header),
		zap.Float64("latency", latency),
	}
	Logger.Info("Request", fields...)
}
