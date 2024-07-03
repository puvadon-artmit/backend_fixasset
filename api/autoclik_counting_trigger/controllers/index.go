package ControllerAutoclik_Counting_Trigger

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Counting_Trigger "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_trigger/services"

	// ControllerAutoclik_data "github.com/puvadon-artmit/gofiber-template/api/autoclik_data/controllers"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Counting_Trigger.GetById(c.Params("autoclik_counting_trigger_id"))
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

func GetByIdAutoclik_Count(c *fiber.Ctx) error {

	autoclikCountingTriggerID := c.Params("autoclik_counting_trigger_id")

	autoclikCountingTrigger, err := ServicesAutoclik_Counting_Trigger.GetByIdAutoclik_CountingArray(autoclikCountingTriggerID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if autoclikCountingTrigger != nil {
		Autoclik_Count_ID := autoclikCountingTrigger.Autoclik_count.Autoclik_countID
		branchAutoclik, err := ServicesAutoclik_Counting_Trigger.GetByBranchIDDB(autoclikCountingTrigger.Autoclik_count.BranchAutoclik_ID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		product_groupAutoclik, err := ServicesAutoclik_Counting_Trigger.GetByAutoclik_countIDDB(Autoclik_Count_ID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Extract required data
		genProdPostingGroup := autoclikCountingTrigger.Autoclik_count.Gen_Prod_Posting_Group

		// Ensure branchAutoclik has at least one element
		if len(branchAutoclik) == 0 {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch data not found",
			})
		}

		branchCode := branchAutoclik[0].Branch_Autoclik.Branch_Code
		if branchCode == "" {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch code is empty",
			})
		}

		var nameProductGroups []string
		for _, pg := range product_groupAutoclik {
			if pg.Name_Product_Group != nil {
				nameProductGroups = append(nameProductGroups, *pg.Name_Product_Group)
			}
		}

		// Call GetByProductGroupHandler with the extracted genProdPostingGroup
		genProdPostingGroupResults, err := ServicesAutoclik_Counting_Trigger.GetByGen_Prod_Posting_GroupDB(genProdPostingGroup)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		productGroups, err := ServicesAutoclik_Counting_Trigger.GetByProductGroupNoDB(nameProductGroups, branchCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Filter Product_Groups based on the condition Item_No == No
		var filteredProductGroups []model.Item_Autoclik_Bin_Code
		for _, pg := range productGroups {
			if pg.Item_No != nil {
				for _, gpg := range genProdPostingGroupResults {
					if gpg.No != nil && *pg.Item_No == *gpg.No {
						filteredProductGroups = append(filteredProductGroups, *pg)
					}
				}
			}
		}

		response := fiber.Map{
			"status": "success",
			"result": fiber.Map{
				"autoclik_count_id":      Autoclik_Count_ID,
				"branch_code":            branchCode,
				"gen_prod_posting_group": genProdPostingGroup,
				"name_product_group":     nameProductGroups,
				"product_group":          filteredProductGroups, // Use filteredProductGroups after filtering
			},
		}

		// Uncomment if you need to call ReaddataAndCreate function
		err = ReaddataAndCreate(filteredProductGroups, Autoclik_Count_ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(response)
	}

	return c.Status(404).JSON(fiber.Map{
		"status": "error",
		"error":  "Data not found",
	})
}

func ReaddataAndCreate(productGroups []model.Item_Autoclik_Bin_Code, autoclikCountID string) error {
	for _, value := range productGroups {
		autoclik_data := model.Autoclik_Count_Store{
			Odata_etag:           value.Odata_etag,
			Location_Code:        value.Location_Code,
			Bin_Code:             value.Bin_Code,
			Item_No:              value.Item_No,
			Variant_Code:         value.Variant_Code,
			Unit_of_Measure_Code: value.Unit_of_Measure_Code,
			ItemDesc:             value.ItemDesc,
			ItemDesc2:            value.ItemDesc2,
			ItemCategory:         value.ItemCategory,
			ProductGroup:         value.ProductGroup,
			Fixed:                value.Fixed,
			Default:              value.Default,
			Dedicated:            value.Dedicated,
			CalcQtyUOM:           value.CalcQtyUOM,
			Quantity_Base:        value.Quantity_Base,
			Bin_Type_Code:        value.Bin_Type_Code,
			Zone_Code:            value.Zone_Code,
			Lot_No_Filter:        value.Lot_No_Filter,
			Serial_No_Filter:     value.Serial_No_Filter,
			Autoclik_countID:     autoclikCountID,
		}

		_, err := ServicesAutoclik_Counting_Trigger.CreateNewAutoclik_Storedata(autoclik_data)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetByProductGroupHandler retrieves details based on a specific gen_prod_posting_group.
func GetByGen_Prod_Posting_GroupHandler(c *fiber.Ctx) error {
	genProdPostingGroup := c.Query("gen_prod_posting_group")
	autoclik, err := ServicesAutoclik_Counting_Trigger.GetByGen_Prod_Posting_GroupDB(genProdPostingGroup)
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
	productGroups := c.Query("product_group")
	locationCode := c.Query("location_code")

	if productGroups == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "product_group query parameter is required",
		})
	}

	if locationCode == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "location_code query parameter is required",
		})
	}

	// Split the query parameter values into a slice
	productGroupList := strings.Split(productGroups, ",")

	err := ServicesAutoclik_Counting_Trigger.ReaddataAndCreate(productGroupList, locationCode)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	autoclik, err := ServicesAutoclik_Counting_Trigger.GetByProductGroupNoDB(productGroupList, locationCode)
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

// func GetByProductGroupHandler(c *fiber.Ctx) error {
// 	productGroups := c.Query("product_group")
// 	locationCode := c.Query("location_code")

// 	if productGroups == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  "product_group query parameter is required",
// 		})
// 	}

// 	if locationCode == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  "location_code query parameter is required",
// 		})
// 	}

// 	// Split the query parameter values into a slice
// 	productGroupList := strings.Split(productGroups, ",")

// 	autoclik, err := ServicesAutoclik_Counting_Trigger.GetByProductGroupNoDB(productGroupList, locationCode)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": autoclik,
// 	})
// }

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesAutoclik_Counting_Trigger.GetAllRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	// ปริ้นออกมาทาง console เมื่อทำงานสำเร็จ
	// fmt.Println("Operation successful!")

	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func Create(c *fiber.Ctx) error {
	autoclik_counting_trigger := new(model.Autoclik_Counting_Trigger)
	err := c.BodyParser(autoclik_counting_trigger)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(autoclik_counting_trigger)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAutoclik_Counting_Trigger.CreateAutoclik_Counting_Trigger(*autoclik_counting_trigger)
	if err != nil {
		if err.Error() == "Autoclik_countID already exists" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdStatus,
	})
}

func UpdateAutoclik_Counting_Trigger(c *fiber.Ctx) error {
	id := c.Params("autoclik_counting_trigger_id")

	updatedType := new(model.Autoclik_Counting_Trigger)
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

	updatedStatus, err := ServicesAutoclik_Counting_Trigger.UpdateAutoclik_Counting_Trigger(id, *updatedType)
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

func DeleteCountingTrigger(c *fiber.Ctx, Autoclik_Counting_TriggerID string) error {
	countingTriggerID := Autoclik_Counting_TriggerID

	err := ServicesAutoclik_Counting_Trigger.DeleteCountingTriggerByID(countingTriggerID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Counting trigger deleted successfully",
	})
}

func GetByAutoclik_Count_ID(c *fiber.Ctx) error {
	Autoclik_Count_ID := c.Params("autoclik_count_id")

	// ดึงข้อมูลและล้างรายการก่อนทำงาน
	// if err := ControllerAutoclik_data.GetAll(c); err != nil {
	// 	return err
	// }

	// if err := ControllerAutoclik_data.Clearandsuckbinitem(c); err != nil {
	// 	return err
	// }

	// รอให้ 2 ฟังชั่นนี้ทำงานเสร็จก่อนทำงาน GetByAutoclik_Count_ID
	return ProcessAutoclikCountID(c, Autoclik_Count_ID)
}

func ProcessAutoclikCountID(c *fiber.Ctx, Autoclik_Count_ID string) error {
	autoclikCountingTrigger, err := ServicesAutoclik_Counting_Trigger.GetByAutoclik_countID(Autoclik_Count_ID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if autoclikCountingTrigger != nil {

		branchAutoclik, err := ServicesAutoclik_Counting_Trigger.GetByBranchIDDB(autoclikCountingTrigger.BranchAutoclik_ID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		product_groupAutoclik, err := ServicesAutoclik_Counting_Trigger.GetByAutoclik_countIDDB(Autoclik_Count_ID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Extract required data
		genProdPostingGroup := autoclikCountingTrigger.Gen_Prod_Posting_Group

		// Ensure branchAutoclik has at least one element
		if len(branchAutoclik) == 0 {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch data not found",
			})
		}

		branchCode := branchAutoclik[0].Branch_Autoclik.Branch_Code
		if branchCode == "" {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch code is empty",
			})
		}

		var nameProductGroups []string
		for _, pg := range product_groupAutoclik {
			if pg.Name_Product_Group != nil {
				nameProductGroups = append(nameProductGroups, *pg.Name_Product_Group)
			}
		}

		// Call GetByProductGroupHandler with the extracted genProdPostingGroup
		genProdPostingGroupResults, err := ServicesAutoclik_Counting_Trigger.GetByGen_Prod_Posting_GroupDB(genProdPostingGroup)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		productGroups, err := ServicesAutoclik_Counting_Trigger.GetByProductGroupNoDB(nameProductGroups, branchCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Filter Product_Groups based on the condition Item_No == No
		var filteredProductGroups []model.Item_Autoclik_Bin_Code
		for _, pg := range productGroups {
			if pg.Item_No != nil {
				for _, gpg := range genProdPostingGroupResults {
					if gpg.No != nil && *pg.Item_No == *gpg.No {
						filteredProductGroups = append(filteredProductGroups, *pg)
					}
				}
			}
		}

		response := fiber.Map{
			"status": "success",
			"result": fiber.Map{
				"autoclik_count_id":      Autoclik_Count_ID,
				"branch_code":            branchCode,
				"gen_prod_posting_group": genProdPostingGroup,
				"name_product_group":     nameProductGroups,
				"product_group":          filteredProductGroups, // Use filteredProductGroups after filtering
			},
		}

		// Uncomment if you need to call ReaddataAndCreate function
		err = ReaddataAndCreate(filteredProductGroups, Autoclik_Count_ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		_, err = ServicesAutoclik_Counting_Trigger.UpdateAutoclikCountStatusPullData(Autoclik_Count_ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(response)
	}

	return c.Status(404).JSON(fiber.Map{
		"status": "error",
		"error":  "Data not found",
	})
}

func UpdateAutoclik_count(c *fiber.Ctx) error {
	id := c.Params("autoclik_count_id")

	updatedCount_autoclik := new(model.Autoclik_count)
	err := c.BodyParser(updatedCount_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedCount_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Update the status_pull_data field for the given id
	updatedStatus, err := ServicesAutoclik_Counting_Trigger.UpdateAutoclikCountStatusPullData(id)
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
