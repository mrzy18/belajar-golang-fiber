package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${time} -> [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-Jan-02",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.All("/log", func(c *fiber.Ctx) error {
		log := map[string]string{
			"log": fmt.Sprintf("%s:%s %d - %s %s", c.IP(), c.Port(), c.Response().StatusCode(), c.Method(), c.Path()),
		}
		return c.JSON(log)
	})
	log.Fatal(app.Listen(":3000"))
}
