package handler

import (
	"log"
	"malai/database"
	"malai/model/emtity"
	"malai/model/request"

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

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	err := c.BodyParser(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	newUser := emtity.UserEmtity{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Age:     user.Age,
	}

	err = database.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to Create",
			"error":   err,
			"docs":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Created successful",
		"docs":    newUser,
		"error":   nil,
	})
}
