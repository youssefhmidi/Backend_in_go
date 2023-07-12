package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewUserRoute(db database.SqliteDatabase, env *bootstrap.Env, ParentRoute *gin.RouterGroup) {
	ul := database.NewUserLogic(db)
	uc := controller.NewUserController(ul, env)

	ParentRoute.GET("/me", uc.Me)
}
