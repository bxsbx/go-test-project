package global

import (
	"context"
	"github.com/gin-gonic/gin"
)

const (
	ctxKey = "context"
)

func GetContext(c *gin.Context) context.Context {
	if v, ok := c.Get(ctxKey); ok {
		if ctx, ok := v.(context.Context); ok {
			return ctx
		}
	}
	return context.Background()
}

func SetContext(c *gin.Context, ctx context.Context) {
	c.Set(ctxKey, ctx)
}
