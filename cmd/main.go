package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/Backend_in_go/bootstrap"
	"github.com/youssefhmidi/Backend_in_go/routes"
)

func main() {
	app := bootstrap.App()

	ParentRoute := gin.Default()
	env := app.Env
	db := app.Database

	routes.SetupRoutes(db, env, ParentRoute)

	ParentRoute.Run()
}
