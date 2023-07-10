package jwtutilities

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/youssefhmidi/Backend_in_go/models"
)

func CreateAccessToken(usr models.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claim := &models.JwtClaims{
		Name: usr.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}

func CreateRefreshToken(usr models.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claim := &models.RefJwtClaims{
		ID: usr.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}

func IsAuthorized(AccesToken string, secret string) (bool, error) {
	_, err := jwt.Parse(AccesToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("cannot identify the signing method used , the signing method used : %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, err
}
