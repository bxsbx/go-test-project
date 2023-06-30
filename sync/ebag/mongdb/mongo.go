package mongodb

import (
	"context"
	beegoConfig "github.com/astaxie/beego/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type mongoDBConfig struct {
	Host       string
	UserName   string
	Password   string
	MaxPool    uint64
	AuthSource string
	DbName     string
}

func defaultMongoDBConfig(cfg beegoConfig.Configer) mongoDBConfig {
	maxPool, _ := cfg.Int("SyncConfig::MongoMaxPool")
	return mongoDBConfig{
		Host:       cfg.String("SyncConfig::MongoHost"),
		UserName:   cfg.String("SyncConfig::MongoUserName"),
		Password:   cfg.String("SyncConfig::MongoPassword"),
		MaxPool:    uint64(maxPool),
		AuthSource: cfg.String("SyncConfig::MongoAuthSource"),
		DbName:     cfg.String("SyncConfig::MongoDbName"),
	}
}

var MongoDB *mongo.Database

func newMongoDB(cfg mongoDBConfig) *mongo.Database {
	clientOps := options.Client().ApplyURI("mongodb://" + cfg.Host).
		SetAuth(options.Credential{
			AuthSource:  cfg.AuthSource,
			Username:    cfg.UserName,
			Password:    cfg.Password,
			PasswordSet: true,
		}).SetMaxPoolSize(cfg.MaxPool)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		log.Fatalf("连接mongodb失败, err:%v", err)
	}
	//检查是否能ping通
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb失败, err:%v", err)
	}

	return client.Database(cfg.DbName)
}

func InitMongoDB(cfg beegoConfig.Configer) {
	MongoDB = newMongoDB(defaultMongoDBConfig(cfg))
}
