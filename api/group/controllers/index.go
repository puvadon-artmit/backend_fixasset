package GroupController

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	GroupServices "github.com/puvadon-artmit/gofiber-template/api/group/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := GroupServices.GetById(c.Params("id"))
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
	value, err := GroupServices.GetAll()
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

func Create(c *fiber.Ctx) error {
	group := new(model.Group)
	err := c.BodyParser(&group)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(group)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdGroup, err := GroupServices.CreateNewGroup(group)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdGroup,
	})
}

func GetCountGroup(c *fiber.Ctx) error {
	count, err := GroupServices.CountLatestRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)
}

func UpdateGroupByID(c *fiber.Ctx) error {
	id := c.Params("group_id")

	updatedType := new(model.Group)
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

	_, err = GroupServices.UpdateGroup(id, *updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedBranch, err := GroupServices.GetById(id)
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

func Deletegroup(c *fiber.Ctx) error {
	GroupID := c.Params("group_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	group, err := GroupServices.GetById(GroupID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = GroupServices.DeleteGroupbyID(group)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": group,
	})
}
