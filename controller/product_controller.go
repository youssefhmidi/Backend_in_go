package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/models"
)

type ProductController struct {
	ProductLogic models.ManipulatorProduct
	Env          *bootstrap.Env
	ShopLogic    models.ManipulatorShop
}

func NewProductController(pl models.ManipulatorProduct, env *bootstrap.Env, sl models.ManipulatorShop) models.ProductRoute {
	return &ProductController{
		ProductLogic: pl,
		Env:          env,
		ShopLogic:    sl,
	}
}

func (pc *ProductController) GetProduct(c *gin.Context) {

}
func (pc *ProductController) AddProduct(c *gin.Context) {

}
