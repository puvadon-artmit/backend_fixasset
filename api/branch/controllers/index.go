package BranchController

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	BranchServices "github.com/puvadon-artmit/gofiber-template/api/branch/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetBybranchId(c *fiber.Ctx) error {
	value, err := BranchServices.GetById(c.Params("branch_id"))
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

func GetAll(c *fiber.Ctx) error {
	value, err := BranchServices.GetAll()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetBranch(c *fiber.Ctx) error {
	count, err := BranchServices.CountBranch()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)
}

func Create(c *fiber.Ctx) error {
	branch := new(model.Branch)
	err := c.BodyParser(&branch)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	group := new(model.Group)
	err = c.BodyParser(&group)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	mainBranch := new(model.Main_branch)
	err = c.BodyParser(&mainBranch)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(branch)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdBranch, err := BranchServices.CreateNewBranch(branch, group, mainBranch)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdBranch, err = BranchServices.GetById(createdBranch.BranchID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdBranch,
	})
}

func Updatebranch(c *fiber.Ctx) error {
	id := c.Params("branch_id")

	updatedType := new(model.Branch)
	err := c.BodyParser(updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	group := new(model.Group)
	err = c.BodyParser(group)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	mainBranch := new(model.Main_branch)
	err = c.BodyParser(mainBranch)
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

	_, err = BranchServices.UpdateBranch(id, *updatedType, group, mainBranch)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedBranch, err := BranchServices.GetById(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedBranch,
	})
}

func Deletebranch(c *fiber.Ctx) error {
	BranchID := c.Params("branch_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	asset_count, err := BranchServices.GetById(BranchID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = BranchServices.DeletebranchbyID(asset_count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": asset_count,
	})
}
