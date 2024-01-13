package main

import "github.com/gofiber/fiber/v2"
import "github.com/whiteyesx/go-practice/database"

func main() {
	app := fiber.New()

	db := database.OpenConnection()
	database.MigrateSchema(db)

	app.Listen(":8080")
}
