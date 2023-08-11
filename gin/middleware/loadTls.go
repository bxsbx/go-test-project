package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 使用 github.com/unrolled/secure 可以方便地为 Go Web 应用程序添加一层安全性保护，避免许多常见的安全漏洞和攻击
// 使用https请求

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		}).Process(c.Writer, c.Request)

		if err != nil {
			return
		}
		c.Next()
	}
}
