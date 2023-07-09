package models

import (
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
	CreateUser(usr *User) error
	GetById(ID uint) (User, error)
	GetByEmail(Email string) (User, error)
}
