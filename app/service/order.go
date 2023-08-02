package service

import (
	"context"
	"ordersystem/ent"
	"ordersystem/internal/pkg/db"
)

type Order struct {
}

func (*Order) Save(data ent.Order, c context.Context) error {
	logger.Info("Order Save")
	order, err := db.Db.Session.Order.Create().SetName(data.Name).Save(c)
	if err != nil {
		logger.Error("create order fail: ", err)
		return err
	}
	logger.Info("create order success", order.String())
	return nil
}
