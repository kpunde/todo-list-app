package routes

import (
	"github.com/gin-gonic/gin"
	"sampleAppDemo/controller"
)

func LoginRoutes(route *gin.Engine) {
	routerGroup := route.Group("/auth")
	{
		routerGroup.POST("/login", controller.Login)
	}
}
