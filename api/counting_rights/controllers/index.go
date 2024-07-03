package ControllerCounting_rights

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesCounting_rights "github.com/puvadon-artmit/gofiber-template/api/counting_rights/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesCounting_rights.GetById(c.Params("counting_rights_id"))
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
	records, err := ServicesCounting_rights.GetAllRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	var result []map[string]interface{}
	for _, record := range records {
		data := map[string]interface{}{
			"counting_rights": record.Counting_rightsID,
			"user_id":         record.UserID,
			"asset_count_id":  record.Asset_countID,
		}
		result = append(result, data)
	}
	return c.JSON(result)
}

func Create(c *fiber.Ctx) error {
	counting_rights := new(model.Counting_rights)
	err := c.BodyParser(counting_rights)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(counting_rights)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesCounting_rights.CreateCounting_rights(*counting_rights)
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

func UpdateCounting_rights(c *fiber.Ctx) error {
	id := c.Params("counting_rights_id")

	updatedType := new(model.Counting_rights)
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

	updatedStatus, err := ServicesCounting_rights.UpdateCounting_rights(id, *updatedType)
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

func DeleteCountingRights(c *fiber.Ctx) error {
	countingRightsID := c.Params("counting_rights_id")

	err := ServicesCounting_rights.DeleteCountingRightsByID(countingRightsID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Counting rights deleted successfully",
	})
}

func GetUSer_IDHandler(c *fiber.Ctx) error {

	value, err := ServicesCounting_rights.GetUserID(c.Params("user_id"))
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
