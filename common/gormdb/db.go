package gormdb

import (
	"fmt"
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dBMap = make(map[string]*gorm.DB)

const (
	DEFAULT = "default"
)

func newDB(cfg dBConfig) *gorm.DB {
	openUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Asia%sShanghai&parseTime=true", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, "%2F")
	db, err := gorm.Open(cfg.DriveDB, openUrl)
	if err != nil {
		log.Fatalf("数据库初始化失败, err:%v", err)
	}
	db.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	db.DB().SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	db.LogMode(cfg.DBLog)
	db.SetLogger(gorm.Logger{LogWriter: log.New(os.Stdout, "\r", 0)})
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
