package handler

import (
	"malai/database"
	"malai/model/emtity"
	"malai/model/request"
	"malai/model/response"

	"github.com/go-playground/validator"
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
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"docs": users,
	})
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user emtity.UserEmtity

	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
			"error":   err,
			"docs":    nil,
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successful",
		"docs":    userResponse,
		"error":   nil,
	})
}

func UserHandlerGetByIdLv2(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user emtity.UserEmtity

	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
			"error":   err,
			"docs":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successful",
		"docs":    user,
		"error":   nil,
	})
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserRequest)

	err := c.BodyParser(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	var validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err,
			"docs":    nil,
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
