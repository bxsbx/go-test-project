package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
	"time"
)

func Co() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		atoi, err := strconv.Atoi("2ew")
		fmt.Println(atoi)
		if err != nil {
			fmt.Println(err)
		}
		return
	})
	c.Start()
	time.Sleep(2 * time.Second)

}

func StartCron() {
	Co()
	time.Sleep(2 * time.Second)
}
