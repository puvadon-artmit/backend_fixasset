package ControllerMaliwan_Fixed_Asset_count_Story

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMaliwan_Fixed_Asset_count_Story "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_count_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Fixed_Asset_count_Story.GetById(c.Params("maliwan_fixed_asset_count_story_id"))
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
	records, err := ServicesMaliwan_Fixed_Asset_count_Story.GetAllRecords()
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
	maliwan_fixed_asset_count_story := new(model.Maliwan_Fixed_Asset_count_Story)
	err := c.BodyParser(maliwan_fixed_asset_count_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(maliwan_fixed_asset_count_story)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_Fixed_Asset_count_Story.CreateNewMaliwan_Fixed_Asset_count_Story(*maliwan_fixed_asset_count_story)
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

func UpdateMaliwan_Fixed_Asset_count_Story(c *fiber.Ctx) error {
	id := c.Params("maliwan_fixed_asset_count_story_id")

	updatedType := new(model.Maliwan_Fixed_Asset_count_Story)
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

	updatedStatus, err := ServicesMaliwan_Fixed_Asset_count_Story.UpdateMaliwan_Fixed_Asset_count_Story(id, *updatedType)
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

func DeleteMaliwan_Fixed_Asset_count_Story(c *fiber.Ctx) error {
	MainBranchStoryID := c.Params("maliwan_fixed_asset_count_story_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	maliwan_fixed_asset_count_story, err := ServicesMaliwan_Fixed_Asset_count_Story.GetById(MainBranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesMaliwan_Fixed_Asset_count_Story.DeleteMaliwan_Fixed_Asset_count_Story(maliwan_fixed_asset_count_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Maliwan_Fixed_Asset_count_Story deleted successfully",
	})
}
