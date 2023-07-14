package database

import (
	"context"

	"github.com/youssefhmidi/Backend_in_go/models"
)

type ProductLogic struct {
	db SqliteDatabase
}

func NewProductLogic(Db SqliteDatabase) models.ManipulatorProduct {
	return &ProductLogic{
		db: Db,
	}
}

func (sl *ProductLogic) AddProducts(ctx context.Context, products []models.Product, shop *models.Shop) error {
	return sl.db.AppendTo("Products", shop, products)
}

func (sl *ProductLogic) GetProducts(ctx context.Context, shop models.Shop, limit int) ([]models.Product, error) {
	Payload, err := sl.db.FindAllByCol(limit, []models.Product{}, "shop_id", shop.ID)
	slice := Payload.([]models.Product)
	return slice, err
}

func (sl *ProductLogic) GetParentShop(ctx context.Context, product models.Product) models.Shop {
	return models.Shop{}
}

func (sl *ProductLogic) FetchAllProducts(ctx context.Context) []models.Product {
	return []models.Product{}
}
