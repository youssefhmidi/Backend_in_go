package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/controller"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func NewSignUpRoute(db database.SqliteDatabase, env *bootstrap.Env, group *gin.RouterGroup) {
	ul := database.NewUserLogic(db)
	sc := controller.NewSignUpController(ul, env)

	group.POST("/signup", sc.SignUp)
}
