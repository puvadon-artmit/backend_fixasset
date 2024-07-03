package ControllerUser_story

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesUser_story "github.com/puvadon-artmit/gofiber-template/api/user_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesUser_story.GetById(c.Params("user_story_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesUser_story.GetAllRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func Create(c *fiber.Ctx) error {
	typeplan := new(model.User_story)
	err := c.BodyParser(typeplan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(typeplan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesUser_story.CreateNewUser_story(*typeplan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdStatus,
	})
}

func UpdateUser_story(c *fiber.Ctx) error {
	id := c.Params("user_story_id")

	updatedType := new(model.User_story)
	err := c.BodyParser(updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedType)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesUser_story.UpdateUser_story(id, *updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedStatus,
	})
}

func DeleteUser_story(c *fiber.Ctx) error {
	typeplanID := c.Params("user_story_id")

	typeplan, err := ServicesUser_story.GetById(typeplanID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesUser_story.DeleteUser_story(typeplan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": typeplan,
	})
}
