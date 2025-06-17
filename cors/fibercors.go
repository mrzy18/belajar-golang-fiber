package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func fiberCors() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://www.google.com",
		AllowMethods: "OPTIONS, GET, POST",
		AllowHeaders: "Content-Type, X-CSRF-Token",
	}))

	app.All("/", func(c *fiber.Ctx) error {
		if c.Method() == "OPTIONS" {
			return c.SendString("Allowed")
		}
		return c.SendString("Hello")
	})

	log.Fatal(app.Listen(":3000"))
}
