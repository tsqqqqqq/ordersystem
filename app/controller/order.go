package controller

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app/service"
	"ordersystem/common"
)

type Order struct {
}

// OrderInfo
// current Order information

func (*Order) Post(c *gin.Context) {
	var order service.OrderInfo
	err := c.BindJSON(&order)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
		return
	}
	err = OrderService.Save(order, c)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
		return
	}
	common.Result(response).Success(c, "create order success", nil)
}
