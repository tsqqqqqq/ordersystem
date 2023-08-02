package controller

import (
	"github.com/gin-gonic/gin"
	"ordersystem/app/service"
	"ordersystem/common"
	"ordersystem/ent"
)

type Inventory struct {
}

var InventoryService service.InventoryService

func (*Inventory) Post(c *gin.Context) {
	var inventory ent.Inventory
	err := c.BindJSON(&inventory)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
	}
	err = InventoryService.Save(inventory, c)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
	}
	common.Result(response).Success(c, "create inventory success", nil)
}
