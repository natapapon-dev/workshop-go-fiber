package routes

import (
	"workshop/go-fiber/routes/api"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	main := app.Group("")
	api.InitApiRoutes(main)
}
