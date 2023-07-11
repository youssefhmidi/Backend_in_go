package bootstrap

import (
	"fmt"

	"github.com/youssefhmidi/Backend_in_go/database"
	"github.com/youssefhmidi/Backend_in_go/models"
)

func InitDB(location string) database.SqliteDatabase {
	DB := database.Database{}
	err := DB.Init(location)

	if err != nil {
		fmt.Println(err.Error())
	}
	DB.Database.AutoMigrate(&models.User{}, &models.Shop{}, &models.Product{})

	return &DB
}
