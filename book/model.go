package book

import (
	"github.com/gofiber/fiber/v2"
	cb "github.com/wcisco17/kubgr/client"
)

func FetchBookById(c *fiber.Ctx) error {
	client, ctx := cb.InitNewClient()
	bookId := c.Params("bookId")
	books, err := FindBook(bookId, client, ctx)
	if err != nil {
		panic(err)
	}

	return c.JSON(books)
}

func FetchBooks(c *fiber.Ctx) error {
	client, ctx := cb.InitNewClient()
	return c.JSON(FindManyBooks(client, ctx))
}
