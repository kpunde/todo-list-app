package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sampleAppDemo/middleware"
	"sampleAppDemo/utility"
)

func HealthCheckRoutes(route *gin.Engine) {
	routerGroup := route.Group("/health", middleware.AuthorizeJWT())
	{
		routerGroup.GET("/overall", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Application is up",
			})
		})

		routerGroup.GET("/database", func(context *gin.Context) {
			utility.Log(zap.DebugLevel, context.GetString("current-user"))
			context.JSON(200, gin.H{
				"message": "Database is up",
			})
		})

	}
}
