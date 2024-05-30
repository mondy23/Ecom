package routes

import (
	"ecom/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(c *fiber.App) {
	c.Get("/", controllers.Home)
}
