package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwtutilities "github.com/youssefhmidi/Backend_in_go/JWT_utilities"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/models"
)

type OrderController struct {
	Env          *bootstrap.Env
	OrdersLogic  models.ManupilatorOrder
	UserLogic    models.ManipulatorUser
	ShopLogic    models.ManipulatorShop
	ProductLogic models.ManipulatorProduct
}

func NewOrderController(env *bootstrap.Env, ol models.ManupilatorOrder, ul models.ManipulatorUser, sl models.ManipulatorShop, pl models.ManipulatorProduct) models.OrdersRoute {
	return &OrderController{
		Env:          env,
		OrdersLogic:  ol,
		UserLogic:    ul,
		ShopLogic:    sl,
		ProductLogic: pl,
	}
}

func (oc *OrderController) PostOrder(c *gin.Context) {
	var req models.OrderReqStructure
	c.ShouldBind(&req)
	access := c.MustGet("Acces_token").(string)

	userId, err := jwtutilities.GetIDFromToken(access, oc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	user, err := oc.UserLogic.GetById(ctx, uint(userId.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	defer cancel()

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	shop, err := oc.ShopLogic.GetShopByID(ctx, req.ShopId)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	Products := []models.Product{}
	for _, v := range req.ProductsId {
		ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
		out, err := oc.ProductLogic.GetProductById(ctx, v)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		}
		Products = append(Products, out)
	}

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	errs := oc.OrdersLogic.PostOrder(ctx, Products, shop, user)
	for _, er := range errs {
		if er != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		}
	}
	defer cancel()

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Order posted"})
}

func (oc *OrderController) GetOrder(c *gin.Context) {

}
