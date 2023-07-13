package database

import (
	"context"

	"github.com/youssefhmidi/Backend_in_go/models"
)

type UserLogic struct {
	db SqliteDatabase
}

func NewUserLogic(Db SqliteDatabase) models.ManipulatorUser {
	return &UserLogic{
		db: Db,
	}
}

func (um *UserLogic) CreateUser(ctx context.Context, usr *models.User) error {
	result := um.db.Add(ctx, usr)

	return result.Error
}
func (um UserLogic) GetById(ctx context.Context, ID uint) (models.User, error) {
	var usr models.User
	um.db.Preload("Shop")
	result := um.db.FindOneById(ctx, &usr, ID)

	return usr, result.Error
}

func (um UserLogic) GetByEmail(ctx context.Context, Email string) (models.User, error) {
	var usr models.User
	um.db.Preload("Shop")
	result := um.db.FindOneByCol(ctx, &usr, "email", Email)

	return usr, result.Error
}
