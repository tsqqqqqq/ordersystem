package service

import (
	"context"
	"errors"
)

type Order struct {
}

type OrderInfo struct {
	InventoryId int64 `json:"inventory_id"` // inventory id
	Count       int   `json:"count"`        // buy count
}

func (*Order) Save(data OrderInfo, c context.Context) error {
	logger.Info("Order Save")
	// TODO First, check if the inventory has enough stock
	var is InventoryService
	// Get current inventory information
	inventory, err, tx := is.GetByIdForTx(data.InventoryId, c)
	if err != nil {
		logger.Error("inventory not found ", err)
		return err
	}
	defer tx.Commit()
	// check if inventory has enough stock
	if inventory.Total < data.Count {
		logger.Error("inventory not enough stock ", err)
		return errors.New("inventory not enough stock")
	}
	logger.Info("update inventory success", inventory.String())
	// create order
	order, err := tx.Order.Create().
		SetInventory(inventory).
		SetName(inventory.Name).
		SetCount(data.Count).
		Save(c)
	if err != nil {
		logger.Error("create order fail: ", err)
		return err
	}
	logger.Info("create order success", order.String())
	// update inventory total
	inventory, err = inventory.Update().SetTotal(inventory.Total - data.Count).Save(c)
	if err != nil {
		logger.Error("update inventory fail: ", err)
		return err
	}
	logger.Info("update inventory success", inventory.String())
	return nil
}
