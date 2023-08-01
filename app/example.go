package app

import (
	"github.com/gin-gonic/gin"
	"ordersystem/common"
)

type Example struct {
}

func (e *Example) Hello(c *gin.Context) {
	response := common.Err
	common.Result(response).Success(c, "hello", nil)
	return
}
