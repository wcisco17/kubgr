package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wcisco17/kubgr/book"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Helo, world!")
	})
	app.Get("/book/c/:title/:desc/:pub", book.CreateBook)
	app.Get("/book/:bookId", book.FetchBookById)
	app.Get("/books", book.FetchBooks)
}

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":8080")
}
