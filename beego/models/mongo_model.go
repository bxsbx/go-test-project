package models

import (
	"StandardProject/common/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PrepareLesson struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	TeacherId   string             `bson:"teacher_id" json:"teacher_id"`
	AreaCode    string             `bson:"area_code" json:"area_code"`
	SchoolId    string             `bson:"school_id,omitempty" json:"school_id,omitempty"`
	CommunityId string             `bson:"community_id,omitempty" json:"community_id,omitempty"`
	BookId      int                `bson:"book_id" json:"book_id"`
	Date        string             `bson:"date" json:"date"`
	From        string             `bson:"from" json:"from"`
	ActionCount int                `bson:"action_count" json:"action_count"`
	UpdateAt    time.Time          `bson:"update_at,omitempty" json:"update_at,omitempty"`
	CreateAt    time.Time          `bson:"create_at,omitempty" json:"create_at,omitempty"`
}

type prepareLessonModel struct {
	db     *mongo.Database
	appCtx context.Context
}

func NewPrepareLessonModel(appCtx context.Context) *prepareLessonModel {
	return &prepareLessonModel{
		db:     mongodb.DefaultMongoDB(),
		appCtx: appCtx,
	}
}
