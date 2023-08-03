package controller

import (
	"ordersystem/app/service"
	"ordersystem/common"
)

var response *common.Response
var InventoryService service.InventoryService
var OrderService *service.Order
