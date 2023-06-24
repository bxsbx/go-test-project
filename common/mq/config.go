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

func defaultMqConfig(cfg beegoConfig.Configer) mqConfig {
	return mqConfig{
		Host:            "",
		Port:            0,
		UserName:        "",
		Password:        "",
		MaxOpenConn:     0,
		MaxIdleConn:     0,
		ConnMaxLifetime: 0,
	}
}
