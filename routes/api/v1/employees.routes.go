package v1

import (
	"workshop/go-fiber/handlers"
	"workshop/go-fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(employees fiber.Router) fiber.Router {
	auth := middlewares.InitAuthBasic()

	employees.Post("/", auth, handlers.CreateEmployee)
	employees.Put("/:id", auth, handlers.UpdateEmployee)
	employees.Delete("/:id", auth, handlers.RemoveEmployeeById)

	employees.Get("/search", handlers.SearchEmployee)
	employees.Get("/generation", handlers.GetEmployeeGeneration)
	employees.Get("/", handlers.GetEmployee)
	employees.Get("/:id", handlers.GetEmployeeById)

	return employees
}
