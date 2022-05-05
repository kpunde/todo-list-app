package routes

import (
	"github.com/gin-gonic/gin"
	"sampleAppDemo/controller"
)

var personController = controller.NewPersonController()

func PersonRoutes(route *gin.Engine) {
	routerGroup := route.Group("/person")
	{
		routerGroup.GET("/:id/items", personController.FindItems)
		routerGroup.GET("/all", personController.FindAll)
		routerGroup.POST("/", personController.Save)
		routerGroup.PUT("/:id", personController.Update)
		routerGroup.DELETE("/:id", personController.Delete)
	}
}
