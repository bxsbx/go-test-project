package main

import (
	_ "TestGeneral/router"
	"github.com/astaxie/beego"
)

//	@title			Swagger Beego
//	@version		1.0
//	@description	这是一个beego框架的swagger测试
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

//go:generate go mod init TestGeneral
//go:generate go mod tidy

//go:generate swag init
func main() {
	beego.Run()
}
