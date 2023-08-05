package main

import (
	"ordersystem/ent"
	"ordersystem/internal/pkg/db"
	"ordersystem/internal/pkg/redis"
	"ordersystem/middleware"
	"ordersystem/routes"
)

func main() {
	logger := middleware.Logger()
	// initialize routes
	logger.Info("Initialize routes...")
	route := routes.Init()

	// use middleware
	middleware.Init(route)

	// initialize db
	_, err := db.NewDb()
	if err != nil {
		return
	}
	defer func(Session *ent.Client) {
		err := Session.Close()
		if err != nil {
			return
		}
	}(db.Db.Session)

	_, err = redis.NewClient()
	if err != nil {
		return
	}

	// Run the server
	logger.Info("Running server...")
	err = route.Run(":8081")
	if err != nil {
		return
	}
}
