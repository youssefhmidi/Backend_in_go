package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwtutilities "github.com/youssefhmidi/Backend_in_go/JWT_utilities"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	UsrLogic models.ManipulatorUser
	Env      *bootstrap.Env
}

func NewLoginController(ul models.ManipulatorUser, env *bootstrap.Env) models.LoginMethod {
	return &LoginController{
		UsrLogic: ul,
		Env:      env,
	}
}

func (lc *LoginController) Login(c *gin.Context) {
	var req models.LoginRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(lc.Env.ContextTimeout))
	usr, err := lc.UsrLogic.GetByEmail(ctx, req.Email)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Wrong Email/user Not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Wrong password"})
		return
	}

	accessToken, err := jwtutilities.CreateAccessToken(usr, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	refreshToken, err := jwtutilities.CreateRefreshToken(usr, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
