package router

import (
	"StandardProject/gin/controller"
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	router := Router.Group("/test")
	api := &controller.Test{}
	router.POST("/json", api.Test1)
	return
}
