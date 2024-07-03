package ControllerCount_autoclik

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Fixed_Asset_Round_Count_Story "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_round_count_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Fixed_Asset_Round_Count_Story.GetById(c.Params("autoclik_fixed_asset_round_count_story_id"))
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
	records, err := ServicesAutoclik_Fixed_Asset_Round_Count_Story.GetAllRecords()
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
	count_autoclik := new(model.Autoclik_Fixed_Asset_Round_Count_Story)
	err := c.BodyParser(count_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(count_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAutoclik_Fixed_Asset_Round_Count_Story.CreateNewAutoclik_Fixed_Asset_Round_Count_Story(*count_autoclik)
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

func UpdateAutoclik_Fixed_Asset_Round_Count_Story(c *fiber.Ctx) error {
	id := c.Params("autoclik_fixed_asset_round_count_story_id")

	updatedCount_autoclik := new(model.Autoclik_Fixed_Asset_Round_Count_Story)
	err := c.BodyParser(updatedCount_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedCount_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesAutoclik_Fixed_Asset_Round_Count_Story.UpdateAutoclik_Fixed_Asset_Round_Count_Story(id, *updatedCount_autoclik)
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

func DeleteAutoclik_Fixed_Asset_Round_Count_Story(c *fiber.Ctx) error {
	BranchStoryID := c.Params("autoclik_fixed_asset_round_count_story_id")

	count_autoclik, err := ServicesAutoclik_Fixed_Asset_Round_Count_Story.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAutoclik_Fixed_Asset_Round_Count_Story.DeleteAutoclik_Fixed_Asset_Round_Count_Story(count_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Count_autoclik deleted successfully",
	})
}

func GetByAutoclik_Fixed_Asset_Round_Count_StoryID(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Fixed_Asset_Round_Count_Story.GetByAutoclik_countIDDB(c.Params("autoclik_fixed_asset_count_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}
