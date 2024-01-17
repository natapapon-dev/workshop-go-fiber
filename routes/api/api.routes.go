package api

import (
	v1 "workshop/go-fiber/routes/api/v1"

	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(main fiber.Router) {
	v1Routes := main.Group("/v1")
	v1.InitV1Routes(v1Routes)
}
