package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InitAuthBasic() fiber.Handler {
	config := basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}

	return basicauth.New(config)
}
