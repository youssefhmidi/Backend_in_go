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
	shopEndpoint := group.Group("/shop")

	NewProductRoute(db, env, shopEndpoint)
	NewOrderRoute(db, env, shopEndpoint)

	shopEndpoint.POST("/create", sc.CreateShop)
	shopEndpoint.GET("/get", sc.GetShop)
	shopEndpoint.GET("/getall", sc.GetAllShop)
	shopEndpoint.PATCH("/edit", sc.EditShop)
	shopEndpoint.DELETE("/delete", sc.DeleteShop)
}
