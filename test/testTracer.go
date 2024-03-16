package main

import (
	"encoding/json"
	"fmt"
)

type DingDingConfig struct {
	OpenRoot  int    `json:"open_root"`  //开启机器人的开关
	RootCode  string `json:"root_code"`  //钉钉机器人code
	AppKey    string `json:"app_key"`    //应用key
	AppSecret string `json:"app_secret"` //应用密钥
}

func main() {
	//cron.StartCron()
	//for i := 0; i < 100; i++ {
	//	time.Sleep(1 * time.Second)
	//}

	jsonS := `{\"open_root\":3,\"root_code\":\"dingrgjuetueiicf8w2y\",\"app_key\":\"dingrgjuetueiicf8w2y\",\"app_secret\":\"I2whwKAV7tOPe1TwKCdamwAmylVv6QGmg0TBh6N54wYSt06sq9fhUvP7lQSF4hPj\"}`
	//jsonS, err := strconv.Unquote(jsonS)
	//if err != nil {
	//	fmt.Println(err)
	//}
	var ding DingDingConfig
	json.Unmarshal([]byte(jsonS), &ding)
	fmt.Println(ding)
}
