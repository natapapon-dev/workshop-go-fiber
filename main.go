package main

import (
	"workshop/go-fiber/config"
	"workshop/go-fiber/database"
	"workshop/go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.InitENV()

	database.InitDatabase()

	routes.InitRoutes(app)

	app.Listen(":3000")
}
