package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewBrowseRoute(db database.SqliteDatabase, env *bootstrap.Env, parentRoute *gin.RouterGroup) {
	sl := database.NewShopLogic(db)

	bc := controller.NewBrowserController(env, sl)
	parentRoute.GET("/browse", bc.Browse)
}
