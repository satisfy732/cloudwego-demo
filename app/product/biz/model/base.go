package model

import "time"

type Base struct {
	ID 		 int `gorm:"primary key"`
	CreateAt time.Time
	UpdatAt  time.Time
}