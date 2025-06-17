package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Data struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	app := fiber.New()

	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/uncompressed"
		},
		Level: compress.LevelBestCompression,
	}))

	app.Get("/compressed", func(c *fiber.Ctx) error {
		data := []Data{}
		for i := 0; i < 1000; i++ {
			data = append(data, Data{
				ID:    i,
				Title: fmt.Sprintf("Title %d", i),
				Body:  fmt.Sprintf("This is a sample body text for item %d", i),
			})
		}
		return c.JSON(data)
	})

	app.Get("/uncompressed", func(c *fiber.Ctx) error {
		data := []Data{}
		for i := 0; i < 1000; i++ {
			data = append(data, Data{
				ID:    i,
				Title: fmt.Sprintf("Title %d", i),
				Body:  fmt.Sprintf("This is a sample body text for item %d", i),
			})
		}
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
