package ControllerBranch_Autoclik

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesBranch_Autoclik "github.com/puvadon-artmit/gofiber-template/api/branch_autoclik/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesBranch_Autoclik.GetById(c.Params("branchautoclik_id"))
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
	records, err := ServicesBranch_Autoclik.GetAllRecords()
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
	branch_autoclik := new(model.Branch_Autoclik)
	err := c.BodyParser(branch_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(branch_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesBranch_Autoclik.CreateNewBranch_Autoclik(*branch_autoclik)
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

func UpdateBranch_Autoclik(c *fiber.Ctx) error {
	id := c.Params("branchautoclik_id")

	updatedType := new(model.Branch_Autoclik)
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

	updatedStatus, err := ServicesBranch_Autoclik.UpdateBranch_Autoclik(id, *updatedType)
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

func DeleteBranch_Autoclik(c *fiber.Ctx) error {
	MainBranchStoryID := c.Params("branchautoclik_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	branch_autoclik, err := ServicesBranch_Autoclik.GetById(MainBranchStoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesBranch_Autoclik.DeleteBranch_Autoclik(branch_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Branch_Autoclik deleted successfully",
	})
}
