package db

import "StandardProject/global"

type Student struct {
	Id    int    `gorm:"column:id;primary_key;"`
	Name  string `gorm:"column:name"`
	Class string `gorm:"column:class"`
	Grade string `gorm:"column:grade"`
}

func (Student) TableName() string {
	return global.StudentTable
}
