package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ShopID      uint
	Name        string
	Description string
	Price       string
	Details     datatypes.JSON
}
