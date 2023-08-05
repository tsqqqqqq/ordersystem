package service

import (
	"context"
	"entgo.io/ent/dialect/sql"
)

type Order struct {
}

type OrderInfo struct {
	InventoryId int64 `json:"inventory_id"` // inventory id
	Count       int   `json:"count"`        // buy count
}

var (
	ctx = context.Background()
)

// Save
// 创建订单信息，保存到数据库。 更新库存信息。
func (*Order) Save(data OrderInfo, c context.Context) error {
	logger.Info("Order Save")
	var inventoryService InventoryService
	// Get current inventory information
	inventory, err, tx := inventoryService.GetByIdForTx(data.InventoryId, c)
	if err != nil {
		logger.Error("inventory not found ", err)
		return err
	}
	// close transaction
	defer tx.Commit()
	// create order
	order, err := tx.Order.Create().
		SetInventory(inventory).
		SetName(inventory.Name).
		SetCount(data.Count).
		Save(c)
	if err != nil {
		tx.Rollback()
		logger.Error("create order fail: ", err)
		return err
	}
	logger.Info("create order success", order.String())
	// TODO 需要修改为订单支付成功后才更新库存
	// update inventory total
	inventory, err = inventory.Update().SetTotal(inventory.Total - data.Count).Where(
		func(selector *sql.Selector) {
			selector.Where(sql.ExprP("total - ? > 0", data.Count))
		},
	).Save(c)
	if err != nil {
		tx.Rollback()
		logger.Error("update inventory fail: ", err)
		return err
	}
	logger.Info("update inventory success", inventory.String())
	return nil
}
