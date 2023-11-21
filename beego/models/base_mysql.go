package models

import (
	"context"
	"gorm.io/gorm"
)

// 对于简短且不常用的sql可使用以下方法，如果sql过长或者复杂或者常用的尽可能地使用具体的方法

type baseMysql struct {
	db     *gorm.DB
	appCtx context.Context
}

func (t *baseMysql) Find(where interface{}, list interface{}, fields ...string) (err error) {
	db := t.db
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	err = db.Where(where).Find(&list).Error
	return
}

func (t *baseMysql) FindWithPageOrder(where interface{}, list interface{}, page, limit int, order string, fields ...string) (err error) {
	db := t.db
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	err = db.Where(where).Order(order).Offset((page - 1) * limit).Limit(limit).Find(&list).Error
	return
}

func (t *baseMysql) First(where interface{}, row interface{}) (err error) {
	err = t.db.Where(where).First(&row).Error
	return
}

func (t *baseMysql) UpdateByWhere(where interface{}, update interface{}) (err error) {
	err = t.db.Where(where).Updates(&update).Error
	return
}

func (t *baseMysql) BatchInsert(list []interface{}) (err error) {
	err = t.db.CreateInBatches(list, 1000).Error
	return
}
