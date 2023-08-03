package routes

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app/controller"
)

var inventoryController *controller.Inventory

func Inventory(route *gin.RouterGroup) {
	route.POST("/inventory", inventoryController.Post)
}
