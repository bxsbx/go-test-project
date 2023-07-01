package main

import (
	_ "StandardProject/sync/ebag/global"
	"StandardProject/sync/ebag/redis"
	"fmt"
)

func main() {
	redis.RedisObj.Set("kok::dd", "csasc")
	str, err := redis.RedisObj.GetString("kok::fr")
	if err != nil {
		//log.Fatal(err)
	}
	//exists, err := redis.RedisObj.Exists("kok::we")
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//err = redis.RedisObj.Remove("kok::dd")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//keys, err := redis.RedisObj.GetKeys("kok::" + "*")
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//fmt.Println(keys)
	//fmt.Println(exists)
	fmt.Println(str)
	fmt.Println("vssv")
}
