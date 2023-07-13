package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewShopRoutes(db database.SqliteDatabase, env *bootstrap.Env, group *gin.RouterGroup) {
	ul := database.NewUserLogic(db)
	sl := database.NewShopLogic(db)
	sc := controller.NewShopController(sl, ul, env)

	group.POST("/shop/create", sc.CreateShop)
	group.GET("/shop/get", sc.GetShop)
	group.GET("/shop/getall", sc.GetAllShop)
	group.PATCH("/shop/edit", sc.EditShop)
	group.DELETE("/shop/delete", sc.DeleteShop)
}
