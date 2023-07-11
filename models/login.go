package models

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginMethod interface {
	Login(c *gin.Context)
}
