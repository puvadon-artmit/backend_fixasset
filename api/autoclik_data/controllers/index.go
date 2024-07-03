package Autoclik_dataController

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	Autoclik_datavices "github.com/puvadon-artmit/gofiber-template/api/autoclik_data/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetAll(c *fiber.Ctx) error {
	err := Autoclik_datavices.ClearAndFetchAllItems()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data retrieved and created successfully",
	})
}

func ClearAndSuck(c *fiber.Ctx) error {
	err := Autoclik_datavices.ClearAndFetchAllItems()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Autoclik retrieved and created successfully",
	})
}

func GetByLimitandOffsetHandler(c *fiber.Ctx) error {
	Limit, _ := strconv.Atoi(c.Query("limit"))
	Offset, _ := strconv.Atoi(c.Query("offset"))

	value, err := Autoclik_datavices.GetAllAutoclikLimit(Limit, Offset)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(value)
}

func Create(c *fiber.Ctx) error {
	autoclik_data := new(model.Item_Autoclik)
	err := c.BodyParser(autoclik_data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(autoclik_data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := Autoclik_datavices.CreateNewAutoclik_data(*autoclik_data)
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

func GetById(c *fiber.Ctx) error {
	value, err := Autoclik_datavices.GetById(c.Params("autoclik_data_id"))
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

func GetByNoHandler(c *fiber.Ctx) error {
	value, err := Autoclik_datavices.GetByNoDB(c.Params("No"))
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

func GetBygen_prod_posting_groupHandler(c *fiber.Ctx) error {
	values, err := Autoclik_datavices.GetBygen_prod_posting_groupDB(c.Params("Gen_Prod_Posting_Group"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": values,
	})
}

func GetAllHandler(c *fiber.Ctx) error {
	autoclik, err := Autoclik_datavices.GetAllAutoclik()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": autoclik,
	})
}

func GetCategory(c *fiber.Ctx) error {
	categories, inventoryGroups, brands, err := Autoclik_datavices.GetAllCategory()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"item_categories":  categories,
		"inventory_groups": inventoryGroups,
		"brands":           brands,
	})
}

func GetByCountHandler(c *fiber.Ctx) error {
	value, err := Autoclik_datavices.GetByCount()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(value)
}
