package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwtutilities "github.com/youssefhmidi/Backend_in_go/JWT_utilities"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/models"
)

type UserController struct {
	UsrLogic models.ManipulatorUser
	Env      *bootstrap.Env
}

func NewUserController(ul models.ManipulatorUser, env *bootstrap.Env) models.UserRoute {
	return &UserController{
		UsrLogic: ul,
		Env:      env,
	}
}
func (uc UserController) Me(c *gin.Context) {
	resp := c.MustGet("Acces_token")
	token := resp.(string)
	Authorized, err := jwtutilities.IsAuthorized(token, uc.Env.AccessTokenSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	if !Authorized {
		c.JSON(http.StatusMethodNotAllowed, models.ErrorResponse{Message: "Access token not verified"})
		return
	}

	out, err := jwtutilities.GetIDFromToken(token, uc.Env.AccessTokenSecret)
	ID := uint(out.(float64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(uc.Env.ContextTimeout))
	usr, err := uc.UsrLogic.GetById(ctx, ID)
	fmt.Println(usr)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}
