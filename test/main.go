package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger_controller.io

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

func GetFileTypeByMimetype(mimeType string, ext ...string) (mimeType2 int) {
	mimeType = strings.ToLower(mimeType)
	if len(ext) <= 0 {
		if strings.Contains(mimeType, "sheet") || strings.Contains(mimeType, "excel") {
			fmt.Println(mimeType)
		}
	} else {
		//extstr := strings.ToLower(ext[0])
		if strings.Contains(mimeType, "excel") {
			fmt.Println(mimeType)
		}
	}
	return
}

//go:generate go run main.go
//go:generate go version
func main() {
	a := "fewef"
	fmt.Println(a[4:])
	atoi, err := strconv.Atoi("08")
	fmt.Println(err, atoi)
}
