package gormdb

import (
	"fmt"
	beegoConfig "github.com/astaxie/beego/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var dBMap = make(map[string]*gorm.DB)

const (
	DEFAULT = "default"
)

func newDB(cfg dBConfig) *gorm.DB {
	openUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Asia%sShanghai&parseTime=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, "%2F")
	db, err := gorm.Open(mysql.Open(openUrl), &gorm.Config{
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
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	return db
}

func InitDB(cfg beegoConfig.Configer) {
	dBMap[DEFAULT] = newDB(defaultDBConfig(cfg))
}

func DefaultDB() *gorm.DB {
	return dBMap[DEFAULT]
}

func GetDB(key string) *gorm.DB {
	return dBMap[key]
}
