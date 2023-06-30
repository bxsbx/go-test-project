package mysql

import (
	"fmt"
	beegoConfig "github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

type dBConfig struct {
	DriveDB         string
	Host            string
	Port            int
	UserName        string
	Password        string
	DbName          string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
	DBLog           bool
}

func defaultDBConfig(cfg beegoConfig.Configer) dBConfig {
	port, _ := cfg.Int("SyncConfig::dbPort")
	maxOpenConn, _ := cfg.Int("SyncConfig::dbMaxConn")
	maxIdleConn, _ := cfg.Int("SyncConfig::maxIdle")
	return dBConfig{
		DriveDB:         "mysql",
		Host:            cfg.String("SyncConfig::dbHost"),
		Port:            port,
		UserName:        cfg.String("SyncConfig::dbUserName"),
		Password:        cfg.String("SyncConfig::dbPassword"),
		DbName:          cfg.String("SyncConfig::dbName"),
		MaxOpenConn:     maxOpenConn,
		MaxIdleConn:     maxIdleConn,
		ConnMaxLifetime: 3600,
		DBLog:           true,
	}
}

var MysqlDB *gorm.DB

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
	MysqlDB = newDB(defaultDBConfig(cfg))
}
