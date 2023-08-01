package routes

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app/controller"
)

var order *controller.Order

func Order(route *gin.RouterGroup) {
	route.POST("/order", order.Post)
}
