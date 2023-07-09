package database

import (
	"github.com/youssefhmidi/Backend_in_go/models"
)

type UserLogic struct {
	db Database
}

func NewUserLogic(Db Database) models.ManipulatorUser {
	return &UserLogic{
		db: Db,
	}
}

func (um *UserLogic) CreateUser(usr *models.User) error {
	result := um.db.Add(&usr)

	return result.Error
}
func (um UserLogic) GetById(ID uint) (models.User, error) {
	var usr models.User
	result := um.db.FindOneByID(&usr, ID)

	return usr, result.Error
}

func (um UserLogic) GetByEmail(Email string) (models.User, error) {
	var usr models.User
	result := um.db.FindOneByCol(usr, "email", Email)

	return usr, result.Error
}
