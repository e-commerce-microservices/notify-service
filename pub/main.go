package main

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello there 👋")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			panic(err)
		}

		return c.SendStatus(200)
	})

	app.Listen(":3000")
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
