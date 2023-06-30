package main

import (
	"StandardProject/sync/ebag/questions"
	"StandardProject/sync/ebag/redis"
)

func main() {
	keys, _ := redis.RedisObj.GetKeys(redis.SQL_UPDATE + "*")
	//keys, _ := redis.RedisObj.GetKeys(redis.SQL_INSERT+"*")
	pageAll := len(keys) / BATCH_SIZE
	for i := 0; i < pageAll; i++ {
		questions.ExecSqlAgain(keys[i*BATCH_SIZE : (i+1)*BATCH_SIZE])
	}
	questions.ExecSqlAgain(keys[pageAll*BATCH_SIZE:])
}
