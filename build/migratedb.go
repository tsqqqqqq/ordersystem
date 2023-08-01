package main

import (
	"context"
	"log"
	"ordersystem/ent"
	"ordersystem/internal/pkg/db"
)

var client *ent.Client

func init() {
	db, err := db.NewDb()
	if err != nil {
		panic(err)
	}
	client = db
}

func main() {
	migrate(context.Background())
	defer client.Close()
}

func migrate(ctx context.Context) {
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
