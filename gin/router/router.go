package router

import (
	"StandardProject/gin/config"
	"StandardProject/gin/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 设置gin模式
	gin.SetMode(config.GetEnvMode())
	var Router = gin.New()

	// 静态文件访问路由（相对地址，文件路径）
	//Router.StaticFS(global.AllConfigInfo.Path, http.Dir(global.AllConfigInfo.Path)) // 为用户头像和文件提供静态地址
	//
	//Router.Use(middleware.LoadTls()) // 该中间间允许使用https请求

	Router.Use(middleware.RecoverPanic()) // 捕获panic并记录(自定义)
	//
	Router.Use(middleware.Common()) // 上下文处理、日志记录
	//
	//Router.Use(middleware.Cors()) // 跨域处理

	//Router.Use(middleware.JWTAuth()) // jwt认证授权

	group := Router.Group("/v1")

	TestRouter(group)
	return Router
}
