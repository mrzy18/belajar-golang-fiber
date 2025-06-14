package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name" form:"name" xml:"name"`
	Email string `json:"email" form:"email" xml:"email"`
}

func main() {
	app := fiber.New()

	// method sendString() untuk render plain text sebagai output
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// method redirect() untuk redirect pengganti http.Redirect()
	app.Get("/redirect", func(c *fiber.Ctx) error {
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	})

	// method JSON() untuk render json sebagai output
	app.Get("/json", func(c *fiber.Ctx) error {
		data := User{
			Name:  "John doe",
			Email: "johndoe@mail.com",
		}
		return c.JSON(data)
	})

	// method Query() untuk mengambil data pada query string request
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		data := fmt.Sprintf("Hello, %s", name)
		return c.SendString(data)
	})

	// method Queries() untuk mengambil lebih dari satu data pada query request
	app.Get("/queries", func(c *fiber.Ctx) error {
		queries := c.Queries()
		data := fmt.Sprintf("Hello, %s. You are %s years old", queries["name"], queries["age"])
		return c.SendString(data)
	})

	// method Params() untuk mengambil data dari path parameter
	app.Get("/params/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		data := fmt.Sprintf("Hello %s", name)
		return c.SendString(data)
	})

	app.All("/user", func(c *fiber.Ctx) error {
		data := new(User)
		if err := c.BodyParser(data); err != nil {
			return err
		}
		return c.JSON(data)
	})

	app.Listen(":3000")
}
