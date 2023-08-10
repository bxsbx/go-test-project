package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
)

func StartCron() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		_, err := strconv.Atoi("2ew")
		//fmt.Println(atoi)
		if err != nil {
			fmt.Println(err)
		}
		return
	})
	c.Start()
}
