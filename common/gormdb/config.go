package gormdb

import (
	beegoConfig "github.com/astaxie/beego/config"
	"log"
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
	port, err := cfg.Int("StandardProject::dbPort")
	if err != nil {
		log.Fatal(err)
	}
	return dBConfig{
		DriveDB:         "mysql",
		Host:            cfg.String("StandardProject::dbHost"),
		Port:            port,
		UserName:        cfg.String("StandardProject::dbUserName"),
		Password:        cfg.String("StandardProject::dbPassword"),
		DbName:          cfg.String("StandardProject::dbName"),
		MaxOpenConn:     20,
		MaxIdleConn:     10,
		ConnMaxLifetime: 3600,
		DBLog:           true,
	}
}
