package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewProductRoute(db database.SqliteDatabase, env *bootstrap.Env, group *gin.RouterGroup) {
	pl := database.NewProductLogic(db)
	sl := database.NewShopLogic(db)
	ul := database.NewUserLogic(db)
	pc := controller.NewProductController(pl, env, sl, ul)

	group.GET("/product", pc.GetProduct)
	group.POST("/product", pc.AddProduct)
}
