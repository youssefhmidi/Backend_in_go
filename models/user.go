package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Shops    []Shop `gorm:"foreignKey:OwnerID"`
}

// interface to create/delete/update the User table
type ManipulatorUser interface {
	CreateUser(ctx context.Context, usr *User) error
	GetById(ctx context.Context, ID uint) (User, error)
	GetByEmail(ctx context.Context, Email string) (User, error)
}
