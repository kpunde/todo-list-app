package routes

import (
	"github.com/gin-gonic/gin"
	"sampleAppDemo/controller"
	"sampleAppDemo/middleware"
)

var itemController = controller.NewItemController()

func ItemRoutes(route *gin.Engine) {
	routerGroup := route.Group("/item", middleware.AuthorizeJWT())
	{

		routerGroup.GET("/all", itemController.FindAll)
		routerGroup.POST("/", itemController.Save)
		routerGroup.PUT("/:id", itemController.Update)
		routerGroup.DELETE("/:id", itemController.Delete)
	}
}
