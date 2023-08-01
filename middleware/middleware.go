package middleware

import "github.com/gin-gonic/gin"

func Init(e *gin.Engine) {
	e.Use(LoggerToFile())
}
