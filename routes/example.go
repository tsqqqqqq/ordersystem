package routes

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app"
)

var example *app.Example

func Example(engine *gin.RouterGroup) {
	engine.GET("/helloworld", example.Hello)
}
