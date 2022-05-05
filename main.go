package main

import (
	"github.com/gin-gonic/gin"
	"sampleAppDemo/middleware"
	"sampleAppDemo/routes"
	"sampleAppDemo/utility"
)

func main() {
	utility.InitDatabase()
	utility.MigrateDB()

	engine := gin.New()

	engine.Use(middleware.Ginzap())
	engine.Use(middleware.RecoveryWithZap(true))

	routes.HealthCheckRoutes(engine)
	routes.ItemRoutes(engine)
	routes.PersonRoutes(engine)
	routes.LoginRoutes(engine)

	engine.Run()
}
