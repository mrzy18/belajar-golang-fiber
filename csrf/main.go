package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html/v2"
)

type User struct {
	Name   string `json:"name" xml:"name" form:"name"`
	Gender string `json:"gender" xml:"gender" form:"gender"`
}

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Use(csrf.New(csrf.Config{
		ContextKey: "token",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		token := c.Locals("token")
		return c.Render("index", fiber.Map{
			"csrf": token,
		})
	})

	app.Post("/sayhello", func(c *fiber.Ctx) error {
		data := &User{}
		if err := c.BodyParser(data); err != nil {
			return err
		}
		message := fmt.Sprintf("Hello %s", data.Name)
		return c.JSON(message)
	})

	log.Fatal(app.Listen(":3000"))
}
