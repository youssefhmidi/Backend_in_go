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

func (pl *ProductLogic) AddProducts(ctx context.Context, products []models.Product, shop *models.Shop) error {
	return pl.db.AppendTo("Products", shop, products, ctx)
}

func (pl *ProductLogic) GetProducts(ctx context.Context, shop models.Shop, limit int) ([]models.Product, error) {
	Payload, err := pl.db.FindAllByCol(limit, []models.Product{}, "shop_id", shop.ID, ctx)
	plice := Payload.([]models.Product)
	return plice, err
}
func (pl *ProductLogic) GetProductById(ctx context.Context, Id uint) (models.Product, error) {
	var out models.Product
	res := pl.db.FindOneById(ctx, &out, Id)
	return out, res.Error
}
