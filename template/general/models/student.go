package models

import (
	//"TestGeneral/common/gormdb"
	"context"
	"gorm.io/gorm"
	"time"
)

const (
	STUDENT_TABLE = "student"
)

type Student struct {
	Id        int       `gorm:"column:id;primary_key"` //主角上次
	Name      string    `gorm:"column:name"`           //dewd
	Class     string    `gorm:"column:class"`          //cece
	Grade     string    `gorm:"column:grade"`          //vwv
	From      int       `gorm:"column:from"`           //茅草删除
	Tr        string    `gorm:"column:tr"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func (Student) TableName() string {
	return STUDENT_TABLE
}

type studentModel struct {
	db     *gorm.DB
	appCtx context.Context
}

func NewStudentModel(appCtx context.Context) *studentModel {
	return &studentModel{
		//db:     gormdb.DefaultDB(),
		appCtx: appCtx,
	}
}

func NewStudentModelWithDB(db *gorm.DB, appCtx context.Context) *studentModel {
	return &studentModel{
		db:     db,
		appCtx: appCtx,
	}
}

func (t *studentModel) Find() (students []Student, err error) {
	err = t.db.Find(&students).Error
	return
}

func (t *studentModel) First(where Student) (student Student, err error) {
	err = t.db.Where(where).First(&student).Error
	return
}

func (t *studentModel) SelectFieldsFindByStudent(fields []string, where Student) (students []Student, err error) {
	if fields == nil {
		err = t.db.Where(where).Find(&students).Error
	} else {
		err = t.db.Select(fields).Where(where).Find(&students).Error
	}
	return
}

func (t *studentModel) FindByPrimary(id int) (students []Student, err error) {
	err = t.db.Where("id = ?", id).Find(&students).Error
	return
}

func (t *studentModel) BatchInsert(list []Student) (err error) {
	err = t.db.CreateInBatches(list, 1000).Error
	return
}

func (t *studentModel) UpdateByWhere(where Student, update Student) (err error) {
	err = t.db.Where(where).Updates(&update).Error
	return
}

func (t *studentModel) DeleteByPrimary(id int) (err error) {
	err = t.db.Where("id = ?", id).Delete(&Student{}).Error
	return
}
