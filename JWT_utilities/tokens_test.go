package jwtutilities

import (
	"log"
	"testing"

	"github.com/youssefhmidi/Backend_in_go/models"
)

func TestCreateAndAuthT(t *testing.T) {
	dummy_usr := models.User{
		Username: "youssef",
	}
	token, err := CreateAccessToken(dummy_usr, "supper_Secret", 1)
	if err != nil {
		log.Fatalf("it got worse, we got this error : %v", err)
	}
	log.Println(token)
	_, Err := IsAuthorized(token, "supper_Secret")
	if Err != nil {
		log.Fatalf("it got worse, we got this error : %v", Err)
	}
}
