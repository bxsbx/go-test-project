package main

import (
	"StandardProject/sync/ebag/questions"
	"StandardProject/sync/ebag/redis"
	"fmt"
)

func main() {
	page := 0
	for {
		key := fmt.Sprintf("%v_%v", MOMENT, page)
		exists, _ := redis.RedisObj.Exists(key)
		if exists {
			page++
			continue
		}
		list, err := questions.GetQsFromBigData(MOMENT, page, BATCH_SIZE)
		if err != nil {
			page++
			continue
		}
		if len(list) <= 0 {
			break
		}
		questions.BatchHandle(list, MOMENT)
		redis.RedisObj.Set(redis.RECORD_SUCCESS_PAGE+key, "ok")
	}
}
