package main

import (
	"StandardProject/common/local"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Map = make(map[string]string)
var lockMap = local.NewLockMap(5, 2)

func M(k string) {
	for i := 0; i < 10; i++ {
		lock := lockMap.GetLock(strconv.Itoa(i))
		lock.Lock()
		fmt.Println(k, i, lock)
		Map[fmt.Sprintf("%v", lock)] = "lock"
		lock.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go M(strconv.Itoa(i))
	}
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	fmt.Println(line, err)
	tem := Map
	tem2 := lockMap
	fmt.Println(tem2)
	fmt.Println(tem)
	//a := &sync.Mutex{}
	//sprintf := fmt.Sprintf("%v", &a)
	//fmt.Println(sprintf)
}
