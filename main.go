package main

import (
	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Helo, world!")
	})
	app.Get("/book/c/:title/:desc/:pub", CreateBook)
	app.Get("/book/:bookId", FetchBookById)
	app.Get("/books", FetchBooks)
}

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":8080")
}
