package controller

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app/service"
	"ordersystem/common"
	"ordersystem/ent"
)

type Order struct {
}

var OrderService service.Order

func (*Order) Post(c *gin.Context) {
	var order ent.Order
	err := c.BindJSON(&order)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
	}
	err = OrderService.Save(order, c)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
	}
	common.Result(response).Success(c, "create order success", nil)
}
