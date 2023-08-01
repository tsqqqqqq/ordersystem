package routes

import (
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() *gin.Engine {
	route := gin.Default()
	v1 := route.Group("/api/v1")
	// 例子
	Example(v1)
	// 订单路由
	Order(v1)
	return route
}
