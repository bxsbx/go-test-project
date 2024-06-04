package middleware

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"strings"
)

// 使用recover（用于从运行时的 panic 中恢复） 从返回值中获取到导致 panic 的具体信息，并用zap做日志记录
// gin中也有自带的
// 只能处理可恢复的panic，对于falatPanic等不可恢复的（程序最终都会终止）
func RecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				//获取所请求的信息
				//httpRequest, _ := httputil.DumpRequest(c.Request, false)

				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				// 检查是连接断开，还是panic(打印错误栈）
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							//zaplog.Logger.Error(c.Request.URL.Path,
							//	zap.Any("error", err),
							//	zap.String("request", string(httpRequest)),
							//)
							// If the connection is dead, we can't write a status to it.
							_ = c.Error(err.(error))
							c.Abort()
							return
						}
					}
				}

				//zaplog.Logger.Error("[Recovery from panic]",
				//	zap.Any("error", err),
				//	zap.String("request", string(httpRequest)),
				//	//zap.String("stack", string(debug.Stack())),
				//)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
