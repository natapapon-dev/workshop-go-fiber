package response

import (
	"workshop/go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func BuidSuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(models.APIResponse{
		APIResponseBase: models.APIResponseBase{
			Data:    data,
			Message: message,
			Success: true,
			Status:  status,
		},
	})
}

func BuidErrorResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(models.APIResponse{
		APIResponseBase: models.APIResponseBase{
			Data:    data,
			Message: message,
			Success: false,
			Status:  status,
		},
	})
}

func BuildSuccessPaginateResponse(c *fiber.Ctx, status int, message string, data interface{}, paginate models.Pagination) error {
	return c.Status(status).JSON(models.APIResponsePaginate{
		APIResponseBase: models.APIResponseBase{
			Data:    data,
			Message: message,
			Success: true,
			Status:  status,
		},
		Pagination: paginate,
	})
}
