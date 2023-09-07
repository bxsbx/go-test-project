package models

import (
	//"TestGeneral/common/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type StudentMongo struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at,omitempty"`
}

type studentMongoModel struct {
	db     *mongo.Database
	appCtx context.Context
}

func NewStudentMongoModel(appCtx context.Context) *studentMongoModel {
	return &studentMongoModel{
		//db:     mongodb.DefaultMongoDB(),
		appCtx: appCtx,
	}
}
