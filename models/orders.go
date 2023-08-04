package models

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ShopID    uint
	OrdererID uint
	Products  []Product `gorm:"many2many:product_order"`
}

type ManupilatorOrder interface {
	PostOrder(ctx context.Context, Products []Product, ParentShop Shop, Orderer User) []error
}

type OrderReqStructure struct {
	ShopId     uint   `json:"shop_id"`
	ProductsId []uint `json:"products"`
}

type OrdersRoute interface {
	PostOrder(c *gin.Context)
	GetOrder(c *gin.Context)
}
