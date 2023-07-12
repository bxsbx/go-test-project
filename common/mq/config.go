package mq

import beegoConfig "github.com/astaxie/beego/config"

type mqConfig struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
}

func defaultMqConfig(cfg beegoConfig.Configer) (config mqConfig) {
	config.Host = ""
	config.Port = 0
	config.UserName = ""
	config.Password = ""
	config.MaxOpenConn = 0
	config.MaxIdleConn = 0
	config.ConnMaxLifetime = 0
	return
}
