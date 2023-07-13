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
	GetShopByID(ctx context.Context, ID uint) (Shop, error)
	GetShopByName(ctx context.Context, Name string) (Shop, error)
	FetchAll(ctx context.Context, limit int) ([]Shop, error)
	UpdateShop(ctx context.Context, Shop *Shop, field string, value interface{}) error
	DeleteShop(ctx context.Context, Shop *Shop) error
}

type ShopRoutes interface {
	GetShop(c *gin.Context)
	CreateShop(c *gin.Context)
	GetAllShop(c *gin.Context)
	EditShop(c *gin.Context)
	DeleteShop(c *gin.Context)
}

type CreateShopRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	IsPrivet bool   `json:"is_privet"`
}

type EditRequest struct {
	ShopID   uint        `json:"id"`
	Field    string      `json:"field"`
	NewValue interface{} `json:"value"`
}
