package main

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

func CreateBookPrisma(client *db.PrismaClient, ctx context.Context, book *Book) *db.BookModel {
	pub := isPub(book)
	books, err := client.Book.CreateOne(
		db.Book.Title.Set(book.Title),
		db.Book.Published.Set(db.Pub(pub)),
		db.Book.Desc.Set(book.Desc),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	return books
}

func isPub(book *Book) string {
	var pubs string = ""
	if book.Published == PUBLISHED {
		pubs = PUBLISHED
	} else {
		pubs = UNPUBLISHED
	}
	return pubs
}
