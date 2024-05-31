package main

import (
	"ecom/pkg/config"
	"ecom/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
