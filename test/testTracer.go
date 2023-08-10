package main

import (
	"StandardProject/cron"
	"time"
)

func main() {
	cron.StartCron()
	for i := 0; i < 100; i++ {
		time.Sleep(1000000 * time.Second)
	}
}
