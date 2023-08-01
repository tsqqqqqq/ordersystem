package main

import (
	"backend-document/ent"
	"backend-document/internal/pkg/db"
	"context"
	"log"
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
