package models

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Shops    []Shop  `gorm:"foreignKey:OwnerID"`
	Orders   []Order `gorm:"foreignKey:OrdererID"`
}

// interface to create/delete/update the User table
type ManipulatorUser interface {
	CreateUser(ctx context.Context, usr *User) error
	GetById(ctx context.Context, ID uint) (User, error)
	GetByEmail(ctx context.Context, Email string) (User, error)
}

type UserRoute interface {
	Me(c *gin.Context)
}
