package main

import (
	_ "StandardProject/sync/ebag/global"
	"StandardProject/sync/ebag/questions"
	"StandardProject/sync/ebag/redis"
	"encoding/json"
	"fmt"
)

func main() {
	//redis.RedisObj.Set("kok::dd", questions.RecordSql{Sql: "sc", Args: []interface{}{"csac", 323, "csc", questions.RecordSql{ExecSql: "vsavasv", Args: []interface{}{"svsvasv", 23}}}})
	str, err := redis.RedisObj.GetString("kok::dd")
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
	var ques questions.RecordSql
	json.Unmarshal([]byte(str), &ques)
	fmt.Println(str)
	fmt.Println("vssv")
}
