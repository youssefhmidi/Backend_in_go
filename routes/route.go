package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func SetupRoutes(db database.SqliteDatabase, env bootstrap.Env, parentRoute *gin.Engine) {
	// public routes
	PublicRoute := parentRoute.Group("")
	NewLoginRoute(db, &env, PublicRoute)
	NewSignUpRoute(db, &env, PublicRoute)

}
