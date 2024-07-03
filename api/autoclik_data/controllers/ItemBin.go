package Autoclik_dataController

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	Autoclik_datavices "github.com/puvadon-artmit/gofiber-template/api/autoclik_data/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetAllBindata(c *fiber.Ctx) error {
	records, err := Autoclik_datavices.GetAllRecords2()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func GetAllGetLocation_Code(c *fiber.Ctx) error {
	records, err := Autoclik_datavices.GetLocation_Code()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func Clearandsuckbinitem(c *fiber.Ctx) error {
	err := Autoclik_datavices.ClearAndFetchAllItemBin()
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

func GetByItem_No_Autoclik_BinHandler(c *fiber.Ctx) error {
	itemNo := c.Query("item_no")
	autoclik, err := Autoclik_datavices.GetByItem_Autoclik_BinNoDB(itemNo)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": autoclik,
	})
}

func GetByProductGroupHandler(c *fiber.Ctx) error {
	product_group := c.Query("product_group")
	autoclik, err := Autoclik_datavices.GetByProductGroupNoDB(product_group)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": autoclik,
	})
}

func GetCountAllItemBin(c *fiber.Ctx) error {
	count, err := Autoclik_datavices.CountAllItemBin()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("%d", count))
}

func GetCountAutoclik_data(c *fiber.Ctx) error {
	count, err := Autoclik_datavices.CountLatestAutoclik_data()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)

}

func GetByLimitandOffsetBindataHandler(c *fiber.Ctx) error {
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

func Suck_information(c *fiber.Ctx) error {
	err := Autoclik_datavices.GetAllRecordsAndCreate2()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "suck information successfully",
	})
}

func CreateBinData(c *fiber.Ctx) error {
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

func GetAllAutoclik_dataHandler(c *fiber.Ctx) error {
	records, err := Autoclik_datavices.GetAllAutoclik_data()
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

func GetBinById(c *fiber.Ctx) error {
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

func GetAllBinHandler(c *fiber.Ctx) error {
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

func GetChoosebytypeAutoclikbin(c *fiber.Ctx) error {
	ItemCategory, ProductGroup, Location_Code, err := Autoclik_datavices.GetChoosebytype()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"ItemCategory":  ItemCategory,
		"ProductGroup":  ProductGroup,
		"Location_Code": Location_Code,
	})
}
