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

type ShopController struct {
	ShopLogic models.ManipulatorShop
	Env       *bootstrap.Env
	UserLogic models.ManipulatorUser
}

func VeryOwnerShip(usr models.User, shop models.Shop) bool {
	return usr.ID == shop.OwnerID
}
func NewShopController(sl models.ManipulatorShop, ul models.ManipulatorUser, env *bootstrap.Env) models.ShopRoutes {
	return &ShopController{
		ShopLogic: sl,
		UserLogic: ul,
		Env:       env,
	}
}

func (sc *ShopController) CreateShop(c *gin.Context) {
	var req models.CreateShopRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	token := c.MustGet("Acces_token").(string)
	Id, err := jwtutilities.GetIDFromToken(token, sc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	ID := uint(Id.(float64))
	user, err := sc.UserLogic.GetById(ctx, ID)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	shop := models.Shop{
		Name:        req.Name,
		Category:    req.Category,
		IsPrivet:    req.IsPrivet,
		Description: req.Description,
	}
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	err = sc.ShopLogic.CreateShop(ctx, shop, &user)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, models.SuccessResponse{Message: "Shop Created"})

}

func (sc *ShopController) GetShop(c *gin.Context) {
	name, IsByName := c.GetQuery("name")
	param, IsById := c.GetQuery("id")

	if IsById {
		uintParm, err := strconv.ParseUint(param, 10, 32)
		Id := uint(uintParm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
		Shop, err := sc.ShopLogic.GetShopByID(ctx, Id)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Shop)
	} else if IsByName {
		ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
		Shop, err := sc.ShopLogic.GetShopByName(ctx, name)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Shop)
		return
	} else {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Wrong/No params are provided"})
		return
	}
}

func (sc *ShopController) GetAllShop(c *gin.Context) {
	token := c.MustGet("Acces_token").(string)
	Raw, err := jwtutilities.GetIDFromToken(token, sc.Env.AccessTokenSecret)
	Id := uint(Raw.(float64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	user, err := sc.UserLogic.GetById(ctx, Id)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, user.Shops)
}

func (sc *ShopController) EditShop(c *gin.Context) {
	var req models.EditRequest
	c.ShouldBind(&req)

	token := c.MustGet("Acces_token").(string)
	Id, err := jwtutilities.GetIDFromToken(token, sc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	ID := uint(Id.(float64))
	user, err := sc.UserLogic.GetById(ctx, ID)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	Shop, err := sc.ShopLogic.GetShopByID(ctx, req.ShopID)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	if !VeryOwnerShip(user, Shop) {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "you don't own this shop by the Id : " + strconv.FormatUint(uint64(req.ShopID), 10)})
		return
	}

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	err = sc.ShopLogic.UpdateShop(ctx, &Shop, req.Field, req.NewValue)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Shop updated"})
}

func (sc *ShopController) DeleteShop(c *gin.Context) {
	type Request struct {
		Id uint `json:"id"`
	}
	var req Request
	c.ShouldBind(&req)
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	shop, err := sc.ShopLogic.GetShopByID(ctx, req.Id)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	token := c.MustGet("Acces_token").(string)
	Id, err := jwtutilities.GetIDFromToken(token, sc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	ID := uint(Id.(float64))
	user, err := sc.UserLogic.GetById(ctx, ID)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	if !VeryOwnerShip(user, shop) {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "you don't own this shop by the Id : " + strconv.FormatUint(uint64(req.Id), 10)})
		return
	}
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	err = sc.ShopLogic.DeleteShop(ctx, &shop)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, models.SuccessResponse{Message: "deleted Shop :" + shop.Name})
}
