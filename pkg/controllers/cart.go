package controllers

import (
	"ecom/pkg/config"
	models "ecom/pkg/models/struct"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Products(c *fiber.Ctx) error {
	var product []models.Product
	db := config.DBconn
	err := db.Find(&product).Error

	if err != nil {
		log.Println("database error")
	}

	return c.JSON(product)
}

func Product(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	db := config.DBconn
	err := db.First(&product, id).Error

	if err != nil {
		log.Println("database error")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "no product found",
		})
	}

	return c.JSON(product)
}

func ViewCart(c *fiber.Ctx) error {
	var order []models.Order
	db := config.DBconn
	db.Joins("Orderdetail").Joins("Orderdetail.Product").Find(&order, 10248)
	return c.JSON(order)
}
