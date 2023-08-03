package service

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"ordersystem/ent"
	inventory2 "ordersystem/ent/inventory"
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

func (InventoryService) GetByIdForTx(id int64, c context.Context) (*ent.Inventory, error, *ent.Tx) {
	logger.Info("Inventory GetById")
	tx, err := db.Db.Session.Tx(c)
	if err != nil {
		logger.Error("start transaction fail: ", err)
	}
	i, err := tx.Inventory.Query().
		Where(inventory2.ID(id)).
		ForUpdate(sql.WithLockAction(sql.NoWait)).
		Only(c)
	if err != nil {
		logger.Error("get inventory fail: ", err)
		return nil, err, tx
	}
	logger.Info("get inventory success", i.String())
	return i, nil, tx
}
