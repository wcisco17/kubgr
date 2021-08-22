package main

import (
	"github.com/gofiber/fiber/v2"
	book "github.com/wcisco17/kubgr/book"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", hello)
	app.Get("/book/:bookId", book.FetchBookById)
	app.Get("/books", book.FetchBooks)
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}
