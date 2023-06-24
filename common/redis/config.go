package redis

import beegoConfig "github.com/astaxie/beego/config"

type redisConfig struct {
	Server      string
	Password    string
	DBNum       int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Wait        bool
}

func defaultRedisConfig(cfg beegoConfig.Configer) redisConfig {
	return redisConfig{
		Server:      "",
		Password:    "",
		DBNum:       1,
		MaxIdle:     23,
		MaxActive:   45,
		IdleTimeout: 2323,
		Wait:        true,
	}
}
