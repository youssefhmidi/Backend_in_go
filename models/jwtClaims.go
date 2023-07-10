package models

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.RegisteredClaims
}

type RefJwtClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}
