package main

import (
	"context"

	"kubgr/db"
)

func InitNewClient() (*db.PrismaClient, context.Context) {
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return nil, nil
	}

	defer func() {
		if err := client.Connect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()
	return client, ctx
}
