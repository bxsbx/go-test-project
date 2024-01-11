package main

import (
	"StandardProject/gin/middleware"
	"StandardProject/swagger/swagger_controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "StandardProject/swagger/docs" //注意路径
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger_controller.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8088

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	//http://127.0.0.1:8088/swagger/index.html访问
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := swagger_controller.NewArticle()

	group1 := r.Group("")
	{
		group1.GET("/api/v1/articles/:id", article.Get)
		group1.PUT("/api/v1/articles", article.Update)
	}

	group2 := r.Group("")
	{
		group2.POST("/api/v1/articles", article.Create)
		group2.DELETE("/api/v1/articles", article.Delete)
	}

	r.Run(":8088")
}
