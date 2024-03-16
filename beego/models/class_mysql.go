package models

const (
	CLASS_TABLE = "class"
)

type Class struct {
	ClassId  string    `gorm:"column:class_id;primary_key;"`
	Name     string    `gorm:"column:name"`
	Students []Student `gorm:"ForeignKey:class_id"`
}

func (Class) TableName() string {
	return CLASS_TABLE
}
