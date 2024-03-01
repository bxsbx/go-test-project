package main

import (
	"time"
)

var timeLocal, _ = time.LoadLocation("Asia/Chongqing")

func main() {
	t := time.Now().Unix()
	d := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	tt, _ := time.Parse("2006-01-02 15:04:05", d)

	println("===========================")
	println("当前时间戳:", t)
	println("当前日期:", d)
	println("从日期得到时间戳:", tt.Unix())
	println("再次转化为日期:", time.Unix(tt.Unix(), 0).Format("2006-01-02 15:04:05"))
	println("===========================")
}
