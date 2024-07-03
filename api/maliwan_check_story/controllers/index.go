package ControllerMaliwan_check_Story

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMaliwan_check_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_check_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_check_Story.GetById(c.Params("maliwan_check_story_id"))
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
	records, err := ServicesMaliwan_check_Story.GetAllRecords()
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
	maliwan_check_story := new(model.Maliwan_check_Story)
	err := c.BodyParser(maliwan_check_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(maliwan_check_story)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_check_Story.CreateNewMaliwan_check_Story(*maliwan_check_story)
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

func UpdateMaliwan_check_Story(c *fiber.Ctx) error {
	id := c.Params("maliwan_check_story_id")

	updatedMaliwan_check_Story := new(model.Maliwan_check_Story)
	err := c.BodyParser(updatedMaliwan_check_Story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedMaliwan_check_Story)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesMaliwan_check_Story.UpdateMaliwan_check_Story(id, *updatedMaliwan_check_Story)
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

func DeleteMaliwan_check_Story(c *fiber.Ctx) error {
	BranchStoryID := c.Params("maliwan_check_story_id")

	maliwan_check_story, err := ServicesMaliwan_check_Story.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesMaliwan_check_Story.DeleteMaliwan_check_Story(maliwan_check_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Maliwan_check_Story deleted successfully",
	})
}
