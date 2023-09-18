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
	db     *gorm.DB
	appCtx context.Context
}

func NewStudentModel(appCtx context.Context) *studentModel {
	return &studentModel{
		db:     gormdb.DefaultDB(),
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

func (t *studentModel) FindByPrimaryKeys(id int, name string) (students []Student, err error) {
	err = t.db.Where("id = ? and name = ?", id, name).Find(&students).Error
	return
}

func (t *studentModel) UpdateByWhere(where Student, update Student) (err error) {
	err = t.db.Where(where).Updates(&update).Error
	return
}

func (t *studentModel) BatchInsert(list []Student) (err error) {
	err = t.db.CreateInBatches(list, 1000).Error
	return
}

func (t *studentModel) DeleteByPrimaryKeys(id int, name string) (err error) {
	err = t.db.Where("id = ? and name = ?", id, name).Delete(&Student{}).Error
	return
}
