package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	pusherClient := pusher.Client{
		AppID:   "1498343",
		Key:     "51f3c1a8e5a382b22ce9",
		Secret:  "f2a71ef86e01df0bfed5",
		Cluster: "eu",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)
		return c.JSON([]string{})
	})

	app.Get("/api/message/get", func(c *fiber.Ctx) error {
		return c.JSON([]string{})
	})

	app.Listen(":8080")
}
