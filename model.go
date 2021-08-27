package main

import "github.com/gofiber/fiber/v2"

func FetchBookById(c *fiber.Ctx) error {
	client, ctx := InitNewClient()
	bookId := c.Params("bookId")
	books, err := FindBook(bookId, client, ctx)
	if err != nil {
		panic(err)
	}

	return c.JSON(books)
}

func FetchBooks(c *fiber.Ctx) error {
	client, ctx := InitNewClient()
	return c.JSON(FindManyBooks(client, ctx))
}

func CreateBook(c *fiber.Ctx) error {
	client, ctx := InitNewClient()

	title := c.Params("title")
	desc := c.Params("desc")
	pub := c.Params("pub")

	return c.JSON(CreateBookPrisma(client, ctx, &Book{
		Title:     title,
		Desc:      desc,
		Published: pub,
	}))
}
