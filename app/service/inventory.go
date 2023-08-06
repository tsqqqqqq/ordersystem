package service

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"ordersystem/ent"
	inventory2 "ordersystem/ent/inventory"
	"ordersystem/internal/pkg/db"
	"ordersystem/internal/pkg/redis"
	"strconv"
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

// SyncInventoryToRedis
// 把库存信息同步到redis
func (InventoryService) SyncInventoryToRedis(ids []int64, c context.Context) error {
	logger.Info("Inventory SyncInventoryToRedis")
	inventorys, err := db.Db.Session.Inventory.Query().Where(inventory2.IDIn(ids...)).All(c)
	if err != nil {
		logger.Error("get inventory fail: ", err)
		return err
	}
	for _, v := range inventorys {
		fmt.Println(v)
		err = redis.Redis_client.Client.HSet(ctx, strconv.FormatInt(v.ID, 10), v).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
