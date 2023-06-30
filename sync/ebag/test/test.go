package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//fmt.Println(fmt.Sprintf(strings.ReplaceAll("?,?", "?", "%v"), "vs", "egge"))
	var wg sync.WaitGroup
	chanNum := make(chan int, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		chanNum <- 1
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			fmt.Println("send success——>", runtime.NumGoroutine())
			<-chanNum
			time.Sleep(1 * time.Second)
			fmt.Println("handle success——>", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
}
