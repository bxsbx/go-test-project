package mongodb

import beegoConfig "github.com/astaxie/beego/config"

type mongoDBConfig struct {
	Host       string
	UserName   string
	Password   string
	MaxPool    uint64
	AuthSource string
	DbName     string
}

func defaultMongoDBConfig(cfg beegoConfig.Configer) (config mongoDBConfig) {
	config.Host = ""
	config.UserName = ""
	config.Password = ""
	config.MaxPool = 12
	config.AuthSource = ""
	config.DbName = ""
	return
}
