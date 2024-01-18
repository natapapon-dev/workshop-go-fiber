package api

import (
	v1 "workshop/go-fiber/routes/api/v1"

	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(main fiber.Router) {
	apiRoutes := main.Group("/api")
	v1.InitV1Routes(apiRoutes)
}
