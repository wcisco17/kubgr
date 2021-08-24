package main

import (
	"github.com/gofiber/fiber/v2"
	book "github.com/wcisco17/kubgr/book"
	cb "github.com/wcisco17/kubgr/client"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Helo, world!")
	})
	app.Get("/book/c/:title/:desc/:pub", func(c *fiber.Ctx) error {
		client, ctx := cb.InitNewClient()

		title := c.Params("title")
		desc := c.Params("desc")
		pub := c.Params("pub")

		return c.JSON(book.CreateBookPrisma(client, ctx, &book.Book{
			Title:     title,
			Desc:      desc,
			Published: pub,
		}))
	})
	app.Get("/book/:bookId", book.FetchBookById)
	app.Get("/books", book.FetchBooks)
}

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}
