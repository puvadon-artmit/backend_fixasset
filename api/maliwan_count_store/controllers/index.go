package ControllerMaliwan_count_Store

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMaliwan_count_Store "github.com/puvadon-artmit/gofiber-template/api/maliwan_count_store/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_count_Store.GetById(c.Params("maliwan_count_store_id"))
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

func GetByMaliwan_Count_StoreIDHandler(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_count_Store.GetByMaliwan_Count_StoreIDDB(c.Params("maliwan_count_id"))
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
	records, err := ServicesMaliwan_count_Store.GetAllRecords()
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
	maliwan_count_store := new(model.Maliwan_Count_Store)
	err := c.BodyParser(maliwan_count_store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(maliwan_count_store)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_count_Store.CreateNewMaliwan_count_Store(*maliwan_count_store)
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

func UpdateMaliwan_count_Store(c *fiber.Ctx) error {
	id := c.Params("maliwan_count_store_id")

	updatedMaliwan_count_Store := new(model.Maliwan_Count_Store)
	err := c.BodyParser(updatedMaliwan_count_Store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedMaliwan_count_Store)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesMaliwan_count_Store.UpdateMaliwan_count_Store(id, *updatedMaliwan_count_Store)
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

func DeleteMaliwan_count_Store(c *fiber.Ctx) error {
	BranchStoryID := c.Params("maliwan_count_store_id")

	maliwan_count_store, err := ServicesMaliwan_count_Store.GetById(BranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesMaliwan_count_Store.DeleteMaliwan_count_Store(maliwan_count_store)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Maliwan_count_Store deleted successfully",
	})
}

func DeleteMaliwan_Count_ID(c *fiber.Ctx) error {
	countingTriggerID := c.Params("maliwan_count_id")

	err := ServicesMaliwan_count_Store.DeleteMaliwan_count_ID(countingTriggerID)
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

func GetByItem_No_Maliwan_BinHandler(c *fiber.Ctx) error {
	itemNo := c.Query("item_no")
	maliwan, err := ServicesMaliwan_count_Store.GetByItem_Maliwan_BinNoDB(itemNo)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": maliwan,
	})
}
