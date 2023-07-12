package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/database"
)

func SetupRoutes(db database.SqliteDatabase, env bootstrap.Env, parentRoute *gin.Engine) {
	// public routes
	PublicRoutes := parentRoute.Group("")
	NewLoginRoute(db, &env, PublicRoutes)
	NewSignUpRoute(db, &env, PublicRoutes)

	// user only routes
	PrivetRoutes := parentRoute.Group("")
	NewUserRoute(db, &env, PrivetRoutes)
}
