package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Todo        string `json:"todo" form:"todo"`
	Description string `json:"description" form:"todo"`
	Status      string `json:"status" form:"status" gorm:"default:not completed"`
	UserID      uint
}
