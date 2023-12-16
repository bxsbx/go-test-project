package main

import (
	"StandardProject/common/dmdb/dmgorm/gorm_v2"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func main() {
	db, err := gorm.Open(gorm_v2.Open("dm://SYSDBA:SYSDBA@127.0.0.1:5236"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库初始化失败, err:%v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Duration(60) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(30) * time.Second)

	var n int64
	db.Table("TEST.CITY").Count(&n)
	fmt.Println(n)
}
