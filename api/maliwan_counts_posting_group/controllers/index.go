package ControllerMaliwan_autoclik

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMaliwan_Counts_Gen_Posting_Group "github.com/puvadon-artmit/gofiber-template/api/maliwan_counts_posting_group/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Counts_Gen_Posting_Group.GetById(c.Params("maliwan_counts_item_category_code_id"))
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
	records, err := ServicesMaliwan_Counts_Gen_Posting_Group.GetAllRecords()
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

func GetMaliwan_count_IDHandler(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Counts_Gen_Posting_Group.GetBymaliwan_count_IDDB(c.Params("maliwan_count_id"))
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
	maliwan_autoclik := new(model.Maliwan_Counts_Item_Category_Code)
	err := c.BodyParser(maliwan_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(maliwan_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_Counts_Gen_Posting_Group.CreateNewMaliwan_Counts_Gen_Posting_Group(*maliwan_autoclik)
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

func UpdateMaliwan_Counts_Gen_Posting_Group(c *fiber.Ctx) error {
	id := c.Params("maliwan_counts_item_category_code_id")

	updatedMaliwan_autoclik := new(model.Maliwan_Counts_Item_Category_Code)
	err := c.BodyParser(updatedMaliwan_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedMaliwan_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesMaliwan_Counts_Gen_Posting_Group.UpdateMaliwan_Counts_Gen_Posting_Group(id, *updatedMaliwan_autoclik)
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

func DeleteMaliwan_Counts_Gen_Posting_Group(c *fiber.Ctx) error {
	BranchStoryID := c.Params("maliwan_counts_item_category_code_id")

	maliwan_autoclik, err := ServicesMaliwan_Counts_Gen_Posting_Group.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesMaliwan_Counts_Gen_Posting_Group.DeleteMaliwan_Counts_Gen_Posting_Group(maliwan_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Maliwan_Counts_Gen_Posting_Group deleted successfully",
	})
}
