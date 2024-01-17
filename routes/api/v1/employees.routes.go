package v1

import (
	"workshop/go-fiber/handlers"
	"workshop/go-fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(employees fiber.Router) fiber.Router {
	auth := middlewares.InitAuthBasic()

	employees.Get("", handlers.GetEmployee)
	employees.Get("/:id", handlers.UpdateEmployee)
	employees.Post("", auth, handlers.CreateEmployee)
	employees.Put("/:id", auth, handlers.CreateEmployee)
	employees.Delete("/:id", auth, handlers.RemoveEmployeeById)
	employees.Get("/search", handlers.SearchEmployee)

	return employees
}
