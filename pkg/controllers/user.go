package controllers

import (
	models "ecom/pkg/models/struct"
	"ecom/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return nil
}

func Register(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	utils.Validate(user.Password, "password")
	return c.JSON(user)
}
