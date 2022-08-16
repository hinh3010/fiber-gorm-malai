package handler

import (
	"log"
	"malai/database"
	"malai/model/emtity"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Message": "Hello, World!",
	})
}

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []emtity.UserEmtity

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	return c.JSON(fiber.Map{
		"docs": users,
	})
}
