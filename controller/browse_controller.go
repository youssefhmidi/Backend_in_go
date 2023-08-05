package controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/models"
)

type BrowserController struct {
	ShopLogic models.ManipulatorShop
	Env       *bootstrap.Env
}

func NewBrowserController(env *bootstrap.Env, sl models.ManipulatorShop) models.BrowseRoutes {
	return &BrowserController{
		Env:       env,
		ShopLogic: sl,
	}
}

func (bc *BrowserController) Browse(c *gin.Context) {
	filters, exist := c.GetQuery("filter")
	filter := strings.Split(filters, "=")
	if !exist {
		ctx, cancel := context.WithTimeout(c, time.Duration(bc.Env.ContextTimeout)*time.Second)
		defer cancel()
		shops, err := bc.ShopLogic.FetchAll(ctx, 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, shops)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Duration(bc.Env.ContextTimeout)*time.Second)
	defer cancel()
	shops, err := bc.ShopLogic.FetchAllByFilter(ctx, 10, filter[0], filter[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	if len(shops) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "cannot find a shop with the provided filter"})
		return
	}
	c.JSON(http.StatusAccepted, shops)
}
