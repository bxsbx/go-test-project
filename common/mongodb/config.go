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

func defaultMongoDBConfig(cfg beegoConfig.Configer) mongoDBConfig {
	return mongoDBConfig{
		Host:       "",
		UserName:   "",
		Password:   "",
		MaxPool:    12,
		AuthSource: "",
		DbName:     "",
	}
}
