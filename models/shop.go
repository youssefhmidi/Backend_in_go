package models

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	OwnerID  uint
	Name     string `gorm:"uniqueIndex"`
	Category string
	IsPrivet bool
	Products []Product `gorm:"ForeignKey:ShopID"`
}

type ManipulatorShop interface {
	CreateShop(ctx context.Context, shop Shop, user *User) error
	AddProducts(ctx context.Context, product []Product) error
	GetShopByID(ctx context.Context, ID uint) (Shop, error)
	GetShopByName(ctx context.Context, Name string) (Shop, error)
	FetchAll(ctx context.Context, limit int) ([]Shop, error)
}

type ShopRoutes interface {
	GetShop(c *gin.Context)
	CreateShop(c *gin.Context)
	GetAllShop(c *gin.Context)
}

type CreateShopRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	IsPrivet bool   `json:"is_privet"`
}
