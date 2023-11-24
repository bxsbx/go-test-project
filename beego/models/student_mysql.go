package models

import (
	"StandardProject/common/gormdb"
	"context"
	"database/sql"
	"gorm.io/gorm"
)

const (
	STUDENT_TABLE = "student"
)

type Student struct {
	Id        int            `gorm:"column:id;primary_key;"`
	Name      string         `gorm:"column:name;primary_key"`
	Class     string         `gorm:"column:class"`
	Grade     sql.NullString `gorm:"column:grade"`
	From      int            `gorm:"column:from"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Student) TableName() string {
	return STUDENT_TABLE
}

type studentModel struct {
	baseMysql
}

func NewStudentModel(appCtx context.Context) *studentModel {
	return &studentModel{
		baseMysql{
			db:     gormdb.DefaultDB(),
			appCtx: appCtx,
		},
	}
}

func NewStudentModelWithDB(db *gorm.DB, appCtx context.Context) *studentModel {
	return &studentModel{
		baseMysql{
			db:     db,
			appCtx: appCtx,
		},
	}
}

type GroupBy struct {
	Name  string `gorm:"column:name"`
	Count int    `gorm:"column:count"`
}

func (t *studentModel) GroupByName(name string) (list []GroupBy, err error) {
	err = t.db.Select("name,count(*) as count").Table("student").Where("name = ?", name).Group("name").Find(&list).Error
	return
}

func (t *studentModel) Test(where Student) (list []Student, err error) {
	err = t.db.Select("name").Table("student").Where(Student{
		Grade: sql.NullString{},
	}).Find(&list).Error
	return
}
