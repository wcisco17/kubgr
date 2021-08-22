package client

import (
	"context"

	"github.com/wcisco17/kubgr/db"
)

func InitNewClient() (*db.PrismaClient, context.Context) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, nil
	}

	defer func() {
		if err := client.Prisma.Connect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()
	return client, ctx
}
