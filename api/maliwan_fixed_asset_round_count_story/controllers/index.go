package ControllerCount_maliwan

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMaliwan_Fixed_Asset_Round_Count_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_round_count_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Fixed_Asset_Round_Count_Story.GetById(c.Params("maliwan_fixed_asset_round_count_story_id"))
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
	records, err := ServicesMaliwan_Fixed_Asset_Round_Count_Story.GetAllRecords()
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
	count_maliwan := new(model.Maliwan_Fixed_Asset_Round_Count_Story)
	err := c.BodyParser(count_maliwan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(count_maliwan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_Fixed_Asset_Round_Count_Story.CreateNewMaliwan_Fixed_Asset_Round_Count_Story(*count_maliwan)
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

func UpdateMaliwan_Fixed_Asset_Round_Count_Story(c *fiber.Ctx) error {
	id := c.Params("maliwan_fixed_asset_round_count_story_id")

	updatedCount_maliwan := new(model.Maliwan_Fixed_Asset_Round_Count_Story)
	err := c.BodyParser(updatedCount_maliwan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedCount_maliwan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesMaliwan_Fixed_Asset_Round_Count_Story.UpdateMaliwan_Fixed_Asset_Round_Count_Story(id, *updatedCount_maliwan)
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

func DeleteMaliwan_Fixed_Asset_Round_Count_Story(c *fiber.Ctx) error {
	BranchStoryID := c.Params("maliwan_fixed_asset_round_count_story_id")

	count_maliwan, err := ServicesMaliwan_Fixed_Asset_Round_Count_Story.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesMaliwan_Fixed_Asset_Round_Count_Story.DeleteMaliwan_Fixed_Asset_Round_Count_Story(count_maliwan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Count_maliwan deleted successfully",
	})
}

func GetByMaliwan_Fixed_Asset_Round_Count_StoryID(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Fixed_Asset_Round_Count_Story.GetByMaliwan_countIDDB(c.Params("maliwan_fixed_asset_count_id"))
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
