package ControllerMain_Category_story

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMain_Category_story "github.com/puvadon-artmit/gofiber-template/api/main_category_story/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMain_Category_story.GetById(c.Params("id"))
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
	records, err := ServicesMain_Category_story.GetAllRecords()
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
	asset_count := new(model.Main_Category_story)
	err := c.BodyParser(asset_count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(asset_count)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMain_Category_story.CreateNewMain_Category_story(*asset_count)
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

func UpdateMain_Category_story(c *fiber.Ctx) error {
	id := c.Params("main_category_story_id")

	updatedType := new(model.Main_Category_story)
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

	updatedStatus, err := ServicesMain_Category_story.UpdateMain_Category_story(id, *updatedType)
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

func DeleteMainCategory(c *fiber.Ctx) error {
	Main_Category_story_ID := c.Params("main_category_story_id")

	err := ServicesMain_Category_story.DeleteMainCategoryByID(Main_Category_story_ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Main_Category_story deleted successfully",
	})
}

func GetByIdHandler(c *fiber.Ctx) error {
	Main_Category_story_ID := c.Params("main_category_story_id")

	mainCategory, err := ServicesMain_Category_story.GetById(Main_Category_story_ID)
	if err != nil {
		return c.Status(404).SendString("")
	}

	return c.JSON(mainCategory)
}
