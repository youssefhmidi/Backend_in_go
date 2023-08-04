package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewOrderRoutes(db database.SqliteDatabase, env *bootstrap.Env, group *gin.RouterGroup) {
	ul, pl, ol, sl := database.NewUserLogic(db),
		database.NewProductLogic(db),
		database.NewOrderLogic(db),
		database.NewShopLogic(db)

	oc := controller.NewOrderController(env, ol, ul, sl, pl)
	group.POST("/order", oc.PostOrder)
	group.POST("/order", oc.GetOrder)
}
