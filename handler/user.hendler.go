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

	// tim data trong db
	err := database.DB.Find(&users).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	// ket qua
	return c.JSON(fiber.Map{
		"docs": users,
	})
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user emtity.UserEmtity

	// tim data trong db
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
			"error":   err,
			"docs":    nil,
		})
	}

	// cac fields data trar ve
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

	// ket qua
	return c.Status(200).JSON(fiber.Map{
		"message": "Successful",
		"docs":    userResponse,
		"error":   nil,
	})
}

func UserHandlerGetByIdLv2(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user emtity.UserEmtity

	// tim data trong db
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"error":   err,
			"docs":    nil,
		})
	}

	// ket qua
	return c.Status(200).JSON(fiber.Map{
		"message": "Successful",
		"docs":    user,
		"error":   nil,
	})
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	// nhan data tu bodyParser
	err := c.BodyParser(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	// validate data
	var validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err,
			"docs":    nil,
		})
	}

	// tao data theo request
	newUser := emtity.UserEmtity{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Age:     user.Age,
	}

	// tao data vao db
	err = database.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to Create",
			"error":   err,
			"docs":    nil,
		})
	}

	// ket qua
	return c.Status(200).JSON(fiber.Map{
		"message": "Created successful",
		"docs":    newUser,
		"error":   nil,
	})
}

func UserHandlerUpdateById(c *fiber.Ctx) error {
	var user emtity.UserEmtity
	userId := c.Params("id")

	// tim data trong db
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"error":   err,
			"docs":    nil,
		})
	}

	// nhan data tu bodyParser
	userRequest := new(request.UserUpdateRequest)
	err = c.BodyParser(userRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	// chinh sua data theo request
	if userRequest.Name != "" && userRequest.Phone != "" {
		user.Name = userRequest.Name
		user.Phone = userRequest.Phone
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name and Phone are not empty",
			"error":   "Bad request",
			"docs":    nil,
		})
	}
	user.Address = userRequest.Address
	user.Age = userRequest.Age

	// luu data vao db
	err = database.DB.Save(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to Update",
			"error":   err,
			"docs":    nil,
		})
	}

	// ket qua
	return c.Status(200).JSON(fiber.Map{
		"message": "Created successful",
		"docs":    user,
		"error":   nil,
	})
}

func UserHandlerDeleteById(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user emtity.UserEmtity

	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"error":   err,
			"docs":    nil,
		})
	}

	err = database.DB.Delete(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to Delete",
			"error":   err,
			"docs":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Deleted successful",
		"error":   nil,
	})
}
