package models

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
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

type ProductRoute interface {
	GetProduct(c *gin.Context)
	AddProduct(c *gin.Context)
}

type ManipulatorProduct interface {
	AddProducts(ctx context.Context, product []Product, shop *Shop) error
	GetProducts(ctx context.Context, shop Shop, limit int) ([]Product, error)
}
