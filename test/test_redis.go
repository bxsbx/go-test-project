package main

import (
	"StandardProject/common/redis"
	_ "StandardProject/global"
	"fmt"
	"sync"
)

func main() {
	obj := redis.DefaultRedisObj()
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			setNX, _ := obj.SetNX("ok", "locked")
			if setNX {
				fmt.Println("get lock")
			}
			mu.Lock()
		}()

	}

}
