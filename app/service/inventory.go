package service

import (
	"context"
	"ordersystem/ent"
	"ordersystem/internal/pkg/db"
)

type InventoryService struct {
}

func (InventoryService) Save(data ent.Inventory, c context.Context) error {
	logger.Info("Inventory Save")
	inventory, err := db.Db.Session.Inventory.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetTotal(data.Total).
		Save(c)
	if err != nil {
		logger.Error("create inventory fail: ", err)
		return err
	}
	logger.Info("create inventory success", inventory.String())
	return nil
}
