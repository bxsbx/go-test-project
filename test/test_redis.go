package main

import (
	"StandardProject/common/redis"
	_ "StandardProject/global"
	"fmt"
)

func main() {
	obj := redis.DefaultRedisObj()
	obj.SetNum("ok", "csc")
	obj.Set("oeek", "csc")
	getString, _ := obj.GetString("ok")
	fmt.Println(getString)
}
