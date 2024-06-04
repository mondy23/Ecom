package routes

import (
	"ecom/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(c *fiber.App) {
	user := c.Group("/user")
	user.Post("/login", controllers.Login)
	user.Post("/register", controllers.Register)

	c.Get("/", controllers.Products)
	c.Post("/:id", controllers.Product)

	cart := c.Group("/cart")
	cart.Get("/", controllers.ViewCart)
}
