package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Println("Hello world")
	app := fiber.New()
	todos := []Todo{}
	app.Post("/api/v1/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(map[string]any{
			"data": todo,
		})
	})
	app.Get("/api/v1/todo", func(c *fiber.Ctx) error {
		return c.JSON(map[string]any{
			"data": todos,
		})
	})
	app.Patch("/api/v1/todo/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id >= len(todos) || id < 0 {
			return c.Status(400).JSON(map[string]string{"message": "invalid todo id provided"})
		}
		todos[id].Done = true
		return c.JSON(map[string]Todo{
			"data": todos[id],
		})
	})

	app.Get("/health-check", func(c *fiber.Ctx) error {
		c.JSON(map[string]string{
			"message": "server is up and running",
		})
		return nil
	})
	app.Use(cors.Config{})
	app.Listen(":8080")
}
