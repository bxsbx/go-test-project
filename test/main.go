package main

import (
	"StandardProject/common/redis"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
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
	//globalLockMap := lock.NewGlobalLockMap()
	//
	//funMap := globalLockMap.GetFunMap("a")
	redis.InitRedis(nil)
	a := 100
	var wg sync.WaitGroup
	for a > 0 {
		wg.Add(1)
		go func(b string, i string) {
			defer wg.Done()
			//funMap := lock.GolaLL.GetFunMap("a")
			//yes := funMap.SetKey(b)
			//if yes {
			//	defer funMap.DelKey(b)
			//	time.Sleep(5 * time.Second)
			//	fmt.Println(b + "---" + i)
			//} else {
			//	fmt.Println("获取" + b + "失败" + "" + i)
			//}

			ok, err := redis.DefaultRedisObj().SetNXTtl(b, "ok", 180)
			if err != nil {
				log.Fatal(err)
			}
			if ok {
				defer redis.DefaultRedisObj().DelKey(b)
				time.Sleep(5 * time.Second)
				fmt.Println(b + "---" + i)
			} else {
				fmt.Println("获取" + b + "失败" + "" + i)
			}
		}(strconv.Itoa(a%2), strconv.Itoa(a))
		a--
	}
	wg.Wait()

}
