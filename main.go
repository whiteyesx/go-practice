package main

import "github.com/gofiber/fiber/v2"

func main() {
	var app = fiber.New()

	app.Listen(":8080")
}
