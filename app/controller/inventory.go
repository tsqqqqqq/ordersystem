package controller

import (
	"github.com/gin-gonic/gin"
	"ordersystem/common"
	"ordersystem/ent"
)

type Inventory struct {
}

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

func (*Inventory) Sync(c *gin.Context) {
	ids := struct {
		Ids []int64 `json:"ids"`
	}{}
	err := c.BindJSON(&ids)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
		return
	}
	err = InventoryService.SyncInventoryToRedis(ids.Ids, c)
	if err != nil {
		common.Result(response).Error(c, err.Error(), nil)
		return
	}
	common.Result(response).Success(c, "sync inventory success", nil)
}
