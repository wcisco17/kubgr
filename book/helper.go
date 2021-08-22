package book

import (
	"context"

	"github.com/wcisco17/kubgr/db"
)

func FindBook(id string, cb *db.PrismaClient, ctx context.Context) (*db.BookModel, error) {
	book, err := cb.Book.FindFirst(
		db.Book.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	return book, nil
}

func FindManyBooks(client *db.PrismaClient, ctx context.Context) []db.BookModel {
	books, err := client.Book.FindMany().Exec(ctx)
	if err != nil {
		panic(err)
	}

	return books
}
