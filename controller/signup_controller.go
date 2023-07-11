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

type SignUpController struct {
	UsrLogic models.ManipulatorUser
	Env      *bootstrap.Env
}

func NewSignUpController(ul models.ManipulatorUser, env *bootstrap.Env) models.SignUpRoute {
	return &SignUpController{
		UsrLogic: ul,
		Env:      env,
	}
}

func (sc SignUpController) SignUp(c *gin.Context) {
	var req models.SignUpRequest

	// checking for request structure
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	// checking if the email is not in use
	ctx, cancel := context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	_, err = sc.UsrLogic.GetByEmail(ctx, req.Email)
	defer cancel()
	if err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "the Email is already used by another user"})
		return
	}

	// generate encrypted password
	pass, Err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if Err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: Err.Error()})
		return
	}

	// create user
	user := models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(pass),
	}
	ctx, cancel = context.WithTimeout(c, time.Second*time.Duration(sc.Env.ContextTimeout))
	err = sc.UsrLogic.CreateUser(ctx, &user)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// creating JWT tokens
	var AccessToken string
	AccessToken, err = jwtutilities.CreateAccessToken(user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	var RefreshToken string
	RefreshToken, err = jwtutilities.CreateRefreshToken(user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// responses to the requester
	c.JSON(http.StatusOK, models.SignUpResponse{AccessToken: AccessToken, RefreshToken: RefreshToken})

}
