package model

import(
	"gorm.io/gorm"
)

type Student struct{
    gorm.Model
	Name    string `gorm:"uniqueIndex;type:varchar(20)"`
	ID string `gorm:"type:varchar(64)"`
	Age int `gorm:"type:int"`
	Class string `gorm:"type:varchar(20)"`
}


