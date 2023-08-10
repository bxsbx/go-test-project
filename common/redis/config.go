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

func defaultRedisConfig(cfg beegoConfig.Configer) (config redisConfig) {
	config.Server = "127.0.0.1"
	config.Password = ""
	config.DBNum = 1
	config.MaxIdle = 23
	config.MaxActive = 45
	config.IdleTimeout = 2323
	config.Wait = true
	return
}
