package routes

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app/controller"
)

var inventoryControler *controller.Inventory

func Inventory(g *gin.RouterGroup) {
	g.POST("/inventory", inventoryControler.Post)
}
