package models

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	OwnerID  uint
	Name     string
	Category string
	IsPrivet bool
	Products []Product `gorm:"ForeignKey:ShopID"`
}
