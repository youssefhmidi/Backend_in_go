package database

import (
	"context"

	"github.com/youssefhmidi/Backend_in_go/models"
)

type OrderLogic struct {
	db SqliteDatabase
}

func NewOrderLogic(Db SqliteDatabase) models.ManupilatorOrder {
	return &OrderLogic{
		db: Db,
	}
}

func (ol *OrderLogic) PostOrder(ctx context.Context, Products []models.Product, ParentShop models.Shop, Orderer models.User) []error {
	order := models.Order{Products: Products}
	res1 := ol.db.AppendTo("Orders", &ParentShop, &order, ctx)
	res2 := ol.db.AppendTo("Orders", &Orderer, &order, ctx)
	res3 := ol.db.AppendTo("Products", &order, Products, ctx)

	if res1 != nil || res2 != nil || res3 != nil {
		return []error{res1, res2, res3}
	}
	return nil
}

func (ol OrderLogic) GetOrderById(ctx context.Context, Id uint) (models.Order, error) {
	var out models.Order
	err := ol.db.FindOneById(ctx, &out, Id).Error
	return out, err
}
