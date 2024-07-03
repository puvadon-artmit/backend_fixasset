package ItemAutoclik_Fixed_AssetController

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	Autoclik_datavices "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetAllAutoclik_Fixed_Asset(c *fiber.Ctx) error {
	records, err := Autoclik_datavices.GetAllAutoclik_Fixed_Asset()
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

func GetAllFA_Location_Code(c *fiber.Ctx) error {
	// ดึงข้อมูล FA_Location_Code และ FA_Posting_Group จาก Autoclik_datavices.GetFA_Location_Code()
	faLocationCodes, faPostingGroups, err := Autoclik_datavices.GetFA_Location_Code()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"fa_location_code": faLocationCodes,
			"fa_posting_group": faPostingGroups,
		},
	})
}

func Clearandsuckbinitem(c *fiber.Ctx) error {
	err := Autoclik_datavices.ClearAllAutoclik_Fixed_Asset()
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
	itemNo := c.Query("no")
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

// func GetByLimitandOffsetBindataHandler(c *fiber.Ctx) error {
// 	Limit, _ := strconv.Atoi(c.Query("limit"))
// 	Offset, _ := strconv.Atoi(c.Query("offset"))

// 	value, err := Autoclik_datavices.GetAllAutoclikLimit(Limit, Offset)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}
// 	return c.JSON(value)
// }

func Pull_information(c *fiber.Ctx) error {
	err := Autoclik_datavices.ClearAndFetchAllAutoclik_Fixed_Asset()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "pull information successfully",
	})
}

func CreateAutoclik_Fixed_AssetData(c *fiber.Ctx) error {
	autoclik_data := new(model.Autoclik_Fixed_Asset)
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

	createdStatus, err := Autoclik_datavices.CreateAutoclik_Fixed_Asset(*autoclik_data)
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
	records, err := Autoclik_datavices.GetAllAutoclik_Fixed_Asset()
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

func GetFilterAutoclik_FixedHandler(c *fiber.Ctx) error {
	FA_Posting_Group := c.Query("fa_posting_group")
	FA_Location_Code := c.Query("fa_location_code")

	if FA_Posting_Group == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "product_group query parameter is required",
		})
	}

	if FA_Location_Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "location_code query parameter is required",
		})
	}

	// Split the query parameter values into a slice
	FA_Posting_Group_List := strings.Split(FA_Posting_Group, ",")

	// err := ServicesAutoclik_Counting_Trigger.ReaddataAndCreate(productGroupList, locationCode)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"status": "error",
	// 		"error":  err.Error(),
	// 	})
	// }

	autoclik, err := Autoclik_datavices.Filter_Data(FA_Posting_Group_List, FA_Location_Code)
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

// func GetBinById(c *fiber.Ctx) error {
// 	value, err := Autoclik_datavices.GetById(c.Params("autoclik_data_id"))
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err,
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": value,
// 	})
// }

// func GetAllBinHandler(c *fiber.Ctx) error {
// 	autoclik, err := Autoclik_datavices.GetAllAutoclik()
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "Not found!",
// 			"error":   err,
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": autoclik,
// 	})
// }

// func GetChoosebytypeAutoclikbin(c *fiber.Ctx) error {
// 	ItemCategory, ProductGroup, Location_Code, err := Autoclik_datavices.GetChoosebytype()
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "Not found!",
// 			"error":   err,
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"ItemCategory":  ItemCategory,
// 		"ProductGroup":  ProductGroup,
// 		"Location_Code": Location_Code,
// 	})
// }

func GetFilterFixedHandler(c *fiber.Ctx) error {

	locationCode := c.Query("locationCode")

	autoclik, err := Autoclik_datavices.GetAutoclik_Fixed_AssetFilter(locationCode)
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
