package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewLoginRoute(db database.SqliteDatabase, env *bootstrap.Env, routeGroup *gin.RouterGroup) {
	ul := database.NewUserLogic(db)
	lc := controller.NewLoginController(ul, env)

	routeGroup.POST("/login", lc.Login)
}
