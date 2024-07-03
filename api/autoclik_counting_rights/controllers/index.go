package ControllerAutoclik_Counting_Rights

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Counting_Rights "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_rights/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Counting_Rights.GetById(c.Params("autoclik_counting_rights_id"))
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
	records, err := ServicesAutoclik_Counting_Rights.GetAllRecords()
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
			"autoclik_counting_rights": record.Autoclik_Counting_RightsID,
			"user_id":                  record.UserID,
			"autoclik_count_id":        record.Autoclik_countID,
		}
		result = append(result, data)
	}
	return c.JSON(result)
}

func Create(c *fiber.Ctx) error {
	autoclik_counting_rights := new(model.Autoclik_Counting_Rights)
	err := c.BodyParser(autoclik_counting_rights)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(autoclik_counting_rights)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAutoclik_Counting_Rights.CreateAutoclik_Counting_Rights(*autoclik_counting_rights)
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

func UpdateAutoclik_Counting_Rights(c *fiber.Ctx) error {
	id := c.Params("autoclik_counting_rights_id")

	updatedType := new(model.Autoclik_Counting_Rights)
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

	updatedStatus, err := ServicesAutoclik_Counting_Rights.UpdateAutoclik_Counting_Rights(id, *updatedType)
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
	countingRightsID := c.Params("autoclik_counting_rights_id")

	err := ServicesAutoclik_Counting_Rights.DeleteCountingRightsByID(countingRightsID)
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

	value, err := ServicesAutoclik_Counting_Rights.GetUserID(c.Params("user_id"))
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
