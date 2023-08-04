package database

import (
	"context"

	"github.com/youssefhmidi/Backend_in_go/models"
)

type ShopLogic struct {
	db   SqliteDatabase
	shop *models.Shop
}

func NewShopLogic(Db SqliteDatabase) models.ManipulatorShop {
	return &ShopLogic{
		db: Db,
	}
}

func (sl *ShopLogic) CreateShop(ctx context.Context, shop models.Shop, user *models.User) error {
	sl.shop = &shop
	return sl.db.AppendTo("Shops", user, []models.Shop{shop}, ctx)
}

func (sl *ShopLogic) GetShopByID(ctx context.Context, ID uint) (models.Shop, error) {
	var shop models.Shop
	res := sl.db.FindOneById(ctx, &shop, ID)
	return shop, res.Error
}

func (sl *ShopLogic) GetShopByName(ctx context.Context, Name string) (models.Shop, error) {
	var shop models.Shop
	res := sl.db.FindOneByCol(ctx, &shop, "name", Name)
	return shop, res.Error
}
func (sl *ShopLogic) FetchAll(ctx context.Context, limit int) ([]models.Shop, error) {
	Payload, err := sl.db.FindAll(limit, []models.Shop{})
	slice := Payload.([]models.Shop)
	return slice, err
}

func (sl *ShopLogic) UpdateShop(ctx context.Context, Shop *models.Shop, field string, value interface{}) error {
	return sl.db.UpdateRow(ctx, Shop, field, value).Error
}

func (sl *ShopLogic) DeleteShop(ctx context.Context, Shop *models.Shop) error {
	res := sl.db.DeleteRecordById(ctx, &models.Shop{}, Shop.ID).Error
	return res
}
