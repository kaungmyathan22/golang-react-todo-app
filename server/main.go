package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title int    `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()
	app.Get("/health-check", func(c *fiber.Ctx) error {
		c.JSON(map[string]string{
			"message": "server is up and running",
		})
		return nil
	})
	app.Listen(":8080")
}
