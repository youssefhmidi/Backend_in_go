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
	Price       int
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

type ProductJSONStructure struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Details     datatypes.JSON `json:"details"`
}

type AddProductRequest struct {
	ShopName string                 `json:"shop_name"`
	Product  []ProductJSONStructure `json:"products"`
}
