package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwtutilities "github.com/youssefhmidi/Backend_in_go/JWT_utilities"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/models"
)

type ProductController struct {
	ProductLogic models.ManipulatorProduct
	Env          *bootstrap.Env
	ShopLogic    models.ManipulatorShop
	UserLogic    models.ManipulatorUser
}

func NewProductController(pl models.ManipulatorProduct, env *bootstrap.Env, sl models.ManipulatorShop, ul models.ManipulatorUser) models.ProductRoute {
	return &ProductController{
		ProductLogic: pl,
		Env:          env,
		ShopLogic:    sl,
		UserLogic:    ul,
	}
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	shopName, exists := c.GetQuery("shop_name")
	if !exists {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "missing required field 'shop_name' in the query"})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(pc.Env.ContextTimeout))
	shop, err := pc.ShopLogic.GetShopByName(ctx, shopName)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, shop.Products)

}
func (pc *ProductController) AddProduct(c *gin.Context) {
	var req models.AddProductRequest

	c.ShouldBind(&req)

	token := c.MustGet("Acces_token").(string)
	Id, err := jwtutilities.GetIDFromToken(token, pc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(pc.Env.ContextTimeout))
	ID := uint(Id.(float64))
	user, err := pc.UserLogic.GetById(ctx, ID)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(pc.Env.ContextTimeout))
	shop, err := pc.ShopLogic.GetShopByName(ctx, req.ShopName)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	if !VeryOwnerShip(user, shop) {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "you don't own this shop by the Id : " + strconv.FormatUint(uint64(shop.ID), 10)})
		return
	}

	var products []models.Product
	for _, values := range req.Product {
		product := models.Product{
			Name:        values.Name,
			Price:       values.Price,
			Description: values.Description,
			Details:     values.Details,
		}
		products = append(products, product)
	}
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(pc.Env.ContextTimeout))
	err = pc.ProductLogic.AddProducts(ctx, products, &shop)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Products added"})
}
