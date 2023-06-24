package models

import (
	"StandardProject/common/gormdb"
	"StandardProject/types/dbType"
	"context"
	"github.com/jinzhu/gorm"
)

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

func (t *studentModel) FindAll() (stu []dbType.Student, err error) {
	err = t.db.Find(&stu).Error
	return
}
