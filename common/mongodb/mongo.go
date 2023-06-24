package mongodb

import (
	"context"
	beegoConfig "github.com/astaxie/beego/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

const (
	DEFAULT = "default"
)

var mongoDBMap = make(map[string]*mongo.Database)

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
	mongoDBMap[DEFAULT] = newMongoDB(defaultMongoDBConfig(cfg))
}

func DefaultMongoDB() *mongo.Database {
	return mongoDBMap[DEFAULT]
}

func GetMongoDB(key string) *mongo.Database {
	return mongoDBMap[key]
}
