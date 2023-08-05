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
func CompareProductsandShop(s models.Shop, p models.Product) bool {
	return s.ID == p.ShopID
}

func (oc *OrderController) PostOrder(c *gin.Context) {
	var req models.OrderReqStructure
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "request structure invalid, see docs"})
		return
	}
	access := c.MustGet("Acces_token").(string)

	userId, err := jwtutilities.GetIDFromToken(access, oc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	user, err := oc.UserLogic.GetById(ctx, uint(userId.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	defer cancel()

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	shop, err := oc.ShopLogic.GetShopByID(ctx, req.ShopId)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	Products := []models.Product{}
	for _, v := range req.ProductsId {
		ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
		out, err := oc.ProductLogic.GetProductById(ctx, v)
		if !CompareProductsandShop(shop, out) {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "this shop don't have some productsin your order"})
			return
		}
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}
		Products = append(Products, out)
	}

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	errs := oc.OrdersLogic.PostOrder(ctx, Products, shop, user)
	defer cancel()
	for _, er := range errs {
		if er != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Order posted"})
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	token := c.MustGet("Acces_token").(string)
	shopId, exist := c.GetQuery("shop_id")
	if !exist {
		userId, err := jwtutilities.GetIDFromToken(token, oc.Env.AccessTokenSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
		user, err := oc.UserLogic.GetById(ctx, uint(userId.(float64)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}
		defer cancel()
		var out []models.Order
		for _, v := range user.Orders {
			ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
			defer cancel()
			res, err := oc.OrdersLogic.GetOrderById(ctx, v.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
				return
			}
			out = append(out, res)
		}

		c.JSON(http.StatusOK, out)
		return
	}
	Id64, err := strconv.ParseUint(shopId, 10, 32)
	Id := uint(Id64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(oc.Env.ContextTimeout))
	shop, err := oc.ShopLogic.GetShopByID(ctx, Id)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, shop.Orders)
}
