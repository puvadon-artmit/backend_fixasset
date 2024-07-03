package ControllerCount_autoclik

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Count_Product_Group "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_posting/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Count_Product_Group.GetById(c.Params("autoclik_count_product_group_id"))
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
	records, err := ServicesAutoclik_Count_Product_Group.GetAllRecords()
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

func GetAutoclik_countHandler(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Count_Product_Group.GetByAutoclik_countIDDB(c.Params("autoclik_count_id"))
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

func Create(c *fiber.Ctx) error {
	count_autoclik := new(model.Autoclik_Count_Product_Group)
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

	createdStatus, err := ServicesAutoclik_Count_Product_Group.CreateNewAutoclik_Count_Product_Group(*count_autoclik)
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

func UpdateAutoclik_Count_Product_Group(c *fiber.Ctx) error {
	id := c.Params("autoclik_count_product_group_id")

	updatedCount_autoclik := new(model.Autoclik_Count_Product_Group)
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

	updatedStatus, err := ServicesAutoclik_Count_Product_Group.UpdateAutoclik_Count_Product_Group(id, *updatedCount_autoclik)
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

func DeleteAutoclik_Count_Product_Group(c *fiber.Ctx) error {
	BranchStoryID := c.Params("autoclik_count_product_group_id")

	count_autoclik, err := ServicesAutoclik_Count_Product_Group.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAutoclik_Count_Product_Group.DeleteAutoclik_Count_Product_Group(count_autoclik)
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
