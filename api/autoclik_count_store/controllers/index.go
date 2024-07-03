package ControllerAutoclik_count_Store

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_count_Store "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_store/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_count_Store.GetById(c.Params("autoclik_count_store_id"))
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

func GetByAutoclik_Count_StoreIDHandler(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_count_Store.GetByAutoclik_Count_StoreIDDB(c.Params("autoclik_count_id"))
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

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesAutoclik_count_Store.GetAllRecords()
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
	autoclik_count_store := new(model.Autoclik_Count_Store)
	err := c.BodyParser(autoclik_count_store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(autoclik_count_store)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAutoclik_count_Store.CreateNewAutoclik_count_Store(*autoclik_count_store)
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

func UpdateAutoclik_count_Store(c *fiber.Ctx) error {
	id := c.Params("autoclik_count_store_id")

	updatedAutoclik_count_Store := new(model.Autoclik_Count_Store)
	err := c.BodyParser(updatedAutoclik_count_Store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedAutoclik_count_Store)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesAutoclik_count_Store.UpdateAutoclik_count_Store(id, *updatedAutoclik_count_Store)
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

func DeleteAutoclik_count_Store(c *fiber.Ctx) error {
	BranchStoryID := c.Params("autoclik_count_store_id")

	autoclik_count_store, err := ServicesAutoclik_count_Store.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAutoclik_count_Store.DeleteAutoclik_count_Store(autoclik_count_store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Autoclik_count_Store deleted successfully",
	})
}

func DeleteAutoclik_Count_ID(c *fiber.Ctx) error {
	countingTriggerID := c.Params("autoclik_count_id")

	err := ServicesAutoclik_count_Store.DeleteAutoclik_count_ID(countingTriggerID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "deleted successfully",
	})
}
