package main

import (
	_ "StandardProject/common/dmdb/dmgorm/gorm_v1"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	db, err := gorm.Open("dm", "dm://SYSDBA:SYSDBA@127.0.0.1:5236")
	if err != nil {
		log.Fatalf("数据库初始化失败, err:%v", err)
	}
	defer db.Close()
	var n int
	db.Table("TEST.CITY").Count(&n)
	fmt.Println(n)
}
