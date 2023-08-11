package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

type T struct {
	F_start_day *time.Time `json:"F_start_day"`
	Num         *int       `json:"num"`
}

//go:generate go run main.go
//go:generate go version
func main() {
	now := time.Now()
	var item T
	item.F_start_day = &now
	c := 23
	item.Num = &c
	marshal, _ := json.Marshal(item)
	fmt.Println(string(marshal))
	var a T
	json.Unmarshal(marshal, &a)
	fmt.Println(*a.Num)
}
