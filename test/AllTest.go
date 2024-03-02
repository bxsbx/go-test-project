package main

import (
	"StandardProject/common/lock"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

var Map = make(map[string]string)
var lockMap = lock.NewLockMap(5, 2)

func M(k string) {
	for i := 0; i < 10; i++ {
		lock := lockMap.GetLock(strconv.Itoa(i))
		lock.Lock()
		fmt.Println(k, i, lock)
		Map[fmt.Sprintf("%v", lock)] = "lock"
		lock.Unlock()
	}
}

type hu struct {
	Str string
}

func main() {
	random, _ := uuid.NewRandom()
	fmt.Println(random.String())
}
