package controllers

import (
	"ecom/pkg/config"
	models "ecom/pkg/models/struct"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var product []models.Product
	db := config.DBconn
	db.Find(&product)
	return c.JSON(product)
}

func Product(c *fiber.Ctx) error {
	return nil
}
