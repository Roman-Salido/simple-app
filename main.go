package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/", helloWorld)

	app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!\n")
}
