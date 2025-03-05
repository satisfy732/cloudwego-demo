package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32
	OrderIdRefer string `gorm:"size:256;index"`
	Quantity     uint32
	Cost         float32
}

func (oi OrderItem) TableName() string {
	return "order_item"
}

