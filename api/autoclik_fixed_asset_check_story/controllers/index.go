package ControllerAutoclik_Fixed_Asset_check_Story

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Fixed_Asset_check_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_check_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Fixed_Asset_check_Story.GetById(c.Params("autoclik_fixed_asset_check_story_id"))
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
	records, err := ServicesAutoclik_Fixed_Asset_check_Story.GetAllRecords()
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
	autoclik_fixed_asset_check_story := new(model.Autoclik_Fixed_Asset_check_Story)
	err := c.BodyParser(autoclik_fixed_asset_check_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(autoclik_fixed_asset_check_story)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAutoclik_Fixed_Asset_check_Story.CreateNewAutoclik_Fixed_Asset_check_Story(*autoclik_fixed_asset_check_story)
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

func UpdateAutoclik_Fixed_Asset_check_Story(c *fiber.Ctx) error {
	id := c.Params("autoclik_fixed_asset_check_story_id")

	updatedAutoclik_Fixed_Asset_check_Story := new(model.Autoclik_Fixed_Asset_check_Story)
	err := c.BodyParser(updatedAutoclik_Fixed_Asset_check_Story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedAutoclik_Fixed_Asset_check_Story)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesAutoclik_Fixed_Asset_check_Story.UpdateAutoclik_Fixed_Asset_check_Story(id, *updatedAutoclik_Fixed_Asset_check_Story)
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

func DeleteAutoclik_Fixed_Asset_check_Story(c *fiber.Ctx) error {
	BranchStoryID := c.Params("autoclik_fixed_asset_check_story_id")

	autoclik_fixed_asset_check_story, err := ServicesAutoclik_Fixed_Asset_check_Story.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAutoclik_Fixed_Asset_check_Story.DeleteAutoclik_Fixed_Asset_check_Story(autoclik_fixed_asset_check_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Autoclik_Fixed_Asset_check_Story deleted successfully",
	})
}
