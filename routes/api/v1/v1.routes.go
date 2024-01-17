package v1

import "github.com/gofiber/fiber/v2"

func InitV1Routes(api fiber.Router) {

	v1 := api.Group("/v1")

	employees := v1.Group("/employees")
	EmployeeRoutes(employees)
}
