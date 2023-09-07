package gormdb

import (
	beegoConfig "github.com/astaxie/beego/config"
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
	ConnMaxIdleTime int
	DBLog           bool
}

func defaultDBConfig(cfg beegoConfig.Configer) (config dBConfig) {
	config.DriveDB = "mysql"
	config.Host = cfg.String("StandardProject::dbHost")
	config.Port, _ = cfg.Int("StandardProject::dbPort")
	config.UserName = cfg.String("StandardProject::dbUserName")
	config.Password = cfg.String("StandardProject::dbPassword")
	config.DbName = cfg.String("StandardProject::dbName")
	config.MaxOpenConn = 20
	config.MaxIdleConn = 10
	config.ConnMaxLifetime = 3600
	config.ConnMaxIdleTime = 3600
	config.DBLog = true
	return
}

func MyDBConfig() (config dBConfig) {
	config.DriveDB = "mysql"
	config.Host = "127.0.0.1"
	config.Port = 3306
	config.UserName = "root"
	config.Password = "123456"
	config.DbName = "test"
	config.MaxOpenConn = 20
	config.MaxIdleConn = 10
	config.ConnMaxLifetime = 3600
	config.ConnMaxIdleTime = 3600
	config.DBLog = true
	return
}
