package ControllerMaliwan_Counting_Trigger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ServicesMaliwan_Counting_Trigger "github.com/puvadon-artmit/gofiber-template/api/maliwan_counting_trigger/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func GetByTrigger_ID(c *fiber.Ctx) error {
	// Fetch Maliwan_Counting_Trigger by its ID
	maliwanCountingTrigger, err := ServicesMaliwan_Counting_Trigger.GetById(c.Params("maliwan_counting_trigger_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Ensure we have a valid Maliwan_Counting_Trigger
	if maliwanCountingTrigger == nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Maliwan_Counting_Trigger not found",
		})
	}

	// Extract Maliwan_countID from the fetched Maliwan_Counting_Trigger
	maliwanCountID := maliwanCountingTrigger.Maliwan_countID

	// Extract Branch_Maliwan from the associated Maliwan_count
	branchMaliwanCount := maliwanCountingTrigger.Maliwan_count.Branch_Maliwan
	if branchMaliwanCount == nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Branch_Maliwan not found",
		})
	}

	Gen_prod_posting_Group_Maliwan := maliwanCountingTrigger.Maliwan_count.Gen_Prod_Posting_Group
	if Gen_prod_posting_Group_Maliwan == "" {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Branch_Maliwan not found",
		})
	}

	// Fetch Item_Category_Code by Maliwan_countID from Maliwan_Counts_Item_Category_Code table
	itemCategoryCodes, err := ServicesMaliwan_Counting_Trigger.GetBymaliwan_count_IDDB(maliwanCountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// No need to loop through itemCategoryCodes to extract Item_Category_Code again
	itemCategoryCodeList := itemCategoryCodes

	// Fetch Maliwan_data by item category codes and branch
	maliwanData, err := ServicesMaliwan_Counting_Trigger.GetByMaliwan_dataNoDB(itemCategoryCodeList, *branchMaliwanCount, Gen_prod_posting_Group_Maliwan)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Prepare the response
	response := fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"maliwan_count_id":         maliwanCountID,
			"maliwan_counting_trigger": maliwanCountingTrigger,
			"item_category_codes":      itemCategoryCodeList,
			"maliwan_data":             maliwanData,
		},
	}

	// err = ReaddataAndCreate(maliwanData, maliwanCountID)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"status": "error",
	// 		"error":  err.Error(),
	// 	})
	// }

	return c.JSON(response)
}

func Create(c *fiber.Ctx) error {
	maliwan_counting_trigger := new(model.Maliwan_Counting_Trigger)
	err := c.BodyParser(maliwan_counting_trigger)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(maliwan_counting_trigger)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_Counting_Trigger.CreateMaliwan_Counting_Trigger(*maliwan_counting_trigger)
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

func UpdateMaliwan_Counting_Trigger(c *fiber.Ctx) error {
	id := c.Params("maliwan_counting_trigger_id")

	updatedType := new(model.Maliwan_Counting_Trigger)
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

	updatedStatus, err := ServicesMaliwan_Counting_Trigger.UpdateMaliwan_Counting_Trigger(id, *updatedType)
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

func DeleteCountingTrigger(c *fiber.Ctx, Maliwan_Counting_TriggerID string) error {
	countingTriggerID := Maliwan_Counting_TriggerID

	err := ServicesMaliwan_Counting_Trigger.DeleteCountingTriggerByID(countingTriggerID)
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

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesMaliwan_Counting_Trigger.GetAllRecords()
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

// -----------------------------------------------------------------------------------------------------------------------------

func ProcessmaliwanCountID(c *fiber.Ctx, maliwanCountID string) error {
	// Fetch Maliwan_Counting_Trigger by its ID
	maliwanCountingTrigger, err := ServicesMaliwan_Counting_Trigger.GetByMaliwan_countID_DB(maliwanCountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Ensure we have a valid Maliwan_Counting_Trigger
	if maliwanCountingTrigger == nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Maliwan_Counting_Trigger not found",
		})
	}

	Gen_prod_posting_Group_Maliwan := maliwanCountingTrigger.Gen_Prod_Posting_Group
	if Gen_prod_posting_Group_Maliwan == "" {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Branch_Maliwan not found",
		})
	}

	// Extract Branch_Maliwan from the associated Maliwan_count
	branchMaliwanCount := maliwanCountingTrigger.Branch_Maliwan
	if branchMaliwanCount == nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Branch_Maliwan not found",
		})
	}

	// Fetch Item_Category_Code by Maliwan_countID from Maliwan_Counts_Item_Category_Code table
	itemCategoryCodes, err := ServicesMaliwan_Counting_Trigger.GetBymaliwan_count_IDDB(maliwanCountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// No need to loop through itemCategoryCodes to extract Item_Category_Code again
	itemCategoryCodeList := itemCategoryCodes

	// Fetch Maliwan_data by item category codes and branch
	maliwanData, err := ServicesMaliwan_Counting_Trigger.GetByMaliwan_dataNoDB(itemCategoryCodeList, *branchMaliwanCount, Gen_prod_posting_Group_Maliwan)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Prepare the response
	response := fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"maliwan_count_id":         maliwanCountID,
			"maliwan_counting_trigger": maliwanCountingTrigger,
			"item_category_codes":      itemCategoryCodeList,
			"maliwan_data":             maliwanData,
		},
	}

	// err = ReaddataAndCreate(maliwanData, maliwanCountID)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"status": "error",
	// 		"error":  err.Error(),
	// 	})
	// }

	return c.JSON(response)
}

// -------------------------------------------   หาค่าจากใบตรวจนับเพื่อไปนำค่าไหา maliwanitembin  ----------------------------------------------------------------------------------

func PullDataByBranch(c *fiber.Ctx, maliwanCountID string) error {
	maliwanCountingTrigger, err := ServicesMaliwan_Counting_Trigger.GetByMaliwan_countID_DB(maliwanCountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if maliwanCountingTrigger == nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Maliwan_Counting_Trigger not found",
		})
	}

	itemCategoryCodes, err := ServicesMaliwan_Counting_Trigger.GetBymaliwan_count_IDDB(maliwanCountID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// No need to loop through itemCategoryCodes to extract Item_Category_Code again
	itemCategoryCodeList := itemCategoryCodes

	Gen_prod_posting_Group_Maliwan := maliwanCountingTrigger.Gen_Prod_Posting_Group
	if Gen_prod_posting_Group_Maliwan == "" {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Branch_Maliwan not found",
		})
	}

	branchMaliwanCount := maliwanCountingTrigger.Branch_Maliwan
	if branchMaliwanCount == nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  "Branch_Maliwan not found",
		})
	}

	// err = ServicesMaliwan_Counting_Trigger.DeleteByBranchMaliwan(*branchMaliwanCount)

	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"status": "error",
	// 		"error":  err.Error(),
	// 	})
	// }

	// สร้าง slice เพื่อเก็บข้อมูล branchMaliwanCount เป็นสมาชิกเดียว
	ArrayofCompanyName := []string{fmt.Sprintf("(%s)", *branchMaliwanCount)}

	response := fmt.Sprintf("(%s)", *branchMaliwanCount)

	err = PullDatamaliwan(c, ArrayofCompanyName, Gen_prod_posting_Group_Maliwan, itemCategoryCodeList, maliwanCountID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	fmt.Println(response)
	return nil
}

func PullDatamaliwan(c *fiber.Ctx, branchMaliwanCounts []string, Gen_prod_posting_Group string, itemCategoryCodes []string, maliwanCountID string) error {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	var records model.Itemapiodata
	var result []model.Itemapivalue

	// วนลูปเพื่อดึงข้อมูลสำหรับทุกสาขา
	for _, branchMaliwanCount := range branchMaliwanCounts {
		fmt.Println(branchMaliwanCount, "ค่าที่รับมา")
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/ItemsAPI"

		// สร้างเงื่อนไขของ filter
		filter := fmt.Sprintf("Gen_Prod_Posting_Group eq '%v'", Gen_prod_posting_Group)
		fmt.Println("Gen_prod_posting_Group = ", Gen_prod_posting_Group)
		var itemCategoryFilter string
		for i, code := range itemCategoryCodes {
			if i > 0 {
				itemCategoryFilter += " or "
			}
			itemCategoryFilter += fmt.Sprintf("Item_Category_Code eq '%v'", code)
		}
		if itemCategoryFilter != "" {
			filter += fmt.Sprintf(" and (%v)", itemCategoryFilter)
			fmt.Println("itemCategoryFilter = ", itemCategoryFilter)
		}

		param := fmt.Sprintf("?$top=%v&$skip=%v&$filter=%v", top, skip, url.QueryEscape(filter))
		path_url2 := NAV_SERVER_HMW + Company + ItemApi + param
		req, err := http.NewRequest("GET", path_url2, strings.NewReader(""))

		if err != nil {
			return err
		}

		for {
			resp, err := Client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
				break
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			err = json.Unmarshal(body, &records)
			if err != nil {
				return err
			}

			result = append(result, records.Value...)

			if len(records.Value) < top {
				break
			}
			skip += top

			param = fmt.Sprintf("?$top=%v&$skip=%v&$filter=%v", top, skip, url.QueryEscape(filter))
			path_url2 = NAV_SERVER_HMW + Company + ItemApi + param
			req, err = http.NewRequest("GET", path_url2, strings.NewReader(""))
			if err != nil {
				return err
			}
		}

		for _, branchMaliwanCount := range branchMaliwanCounts {
			branchName := strings.Trim(branchMaliwanCount, "()")

			var DataFilter []SimplifiedItem
			for _, value := range result {
				item := SimplifiedItem{
					No:          value.No,
					Description: value.Description,
				}
				DataFilter = append(DataFilter, item)
			}

			var NoSlice []string
			for _, value := range result {
				NoSlice = append(NoSlice, value.No)
			}

			// เรียกใช้ฟังก์ชัน GetAllbranchsMaliwan โดยส่งข้อมูลที่จำเป็น
			maliwan, err := ServicesMaliwan_Counting_Trigger.GetAllbranchsMaliwan(branchName, NoSlice)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}

			var maliwanData []model.Maliwan_Count_Store

			for _, maliwanItem := range maliwan {
				// หาค่า Description ที่ตรงกับ Item_No
				var description string
				for _, value := range result {
					if value.No == *maliwanItem.Item_No {
						description = value.Description
						break
					}
				}

				maliwan_data := model.Maliwan_Count_Store{
					Location_Code:        maliwanItem.Location_Code,
					Bin_Code:             maliwanItem.Bin_Code,
					Item_No:              maliwanItem.Item_No,
					Variant_Code:         maliwanItem.Variant_Code,
					Unit_of_Measure_Code: maliwanItem.Unit_of_Measure_Code,
					Fixed:                maliwanItem.Fixed,
					Default:              maliwanItem.Default,
					Dedicated:            maliwanItem.Dedicated,
					CalcQtyUOM:           maliwanItem.CalcQtyUOM,
					Quantity_Base:        maliwanItem.Quantity_Base,
					Bin_Type_Code:        maliwanItem.Bin_Type_Code,
					Zone_Code:            maliwanItem.Zone_Code,
					Lot_No_Filter:        maliwanItem.Lot_No_Filter,
					Serial_No_Filter:     maliwanItem.Serial_No_Filter,
					Description:          &description, // เพิ่ม Description ที่ตรงกับ Item_No
					Branch:               &branchName,
					Maliwan_countID:      maliwanCountID,
				}

				maliwanData = append(maliwanData, maliwan_data)
			}

			// Batch insert
			chunkSize := 1000 // ปรับขนาด chunk ตามความเหมาะสม
			for i := 0; i < len(maliwanData); i += chunkSize {
				end := i + chunkSize
				if end > len(maliwanData) {
					end = len(maliwanData)
				}
				chunk := maliwanData[i:end]

				err := ServicesMaliwan_Counting_Trigger.CreateMaliwan_Count_StoreBatch(chunk)
				if err != nil {
					return err
				}
			}

			_, err = ServicesMaliwan_Counting_Trigger.UpdateMaliwanCountStatusPullData(maliwanCountID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}

			fmt.Println(len(maliwan))
			return c.JSON(maliwan)
		}
	}
	return nil
}

type SimplifiedItem struct {
	No          string `json:"No"`
	Description string `json:"Description"`
}

func GetByMaliwan_CountID(c *fiber.Ctx) error {
	maliwanCountID := c.Params("maliwan_count_id")

	if err := PullDataByBranch(c, maliwanCountID); err != nil {
		return err
	}

	return nil

	// รอให้ PullDataByBranch ฟังชั่นนี้ทำงานเสร็จก่อนทำงาน ProcessmaliwanCountID
	// return ProcessmaliwanCountID(c, maliwanCountID)
}

// ---------------------------------------------  เพิ่มข้อมูลเข้าแต่ล่ะสาขา  --------------------------------------------------------------------

func PullItembin(c *fiber.Ctx) ([]model.Item_bin_maliwan, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := [...]string{"(Head Office)", "(Mueang Krabi)", "(Naka)", "(Burirum)", "(Surin)"}
	var allResults []model.Item_bin_maliwan

	for _, branchMaliwanCount := range branchMaliwanCounts {
		var result []model.Item_bin_maliwan

		fmt.Println(branchMaliwanCount, "ค่าที่รับมา")
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/ItemBinCodeAPI"

		for {
			param := fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
			pathURL := NAV_SERVER_HMW + Company + ItemApi + param
			req, err := http.NewRequest("GET", pathURL, nil)
			if err != nil {
				return nil, err
			}

			resp, err := Client.Do(req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
				break
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			var records model.Itemabinmaliwanodata
			err = json.Unmarshal(body, &records)
			if err != nil {
				return nil, err
			}

			result = append(result, records.Value...)

			if len(records.Value) < top {
				break
			}
			skip += top
		}

		for _, value := range result {
			branchName := strings.Trim(branchMaliwanCount, "()")

			maliwan_data := model.Item_bin_maliwan{
				Location_Code:        value.Location_Code,
				Bin_Code:             value.Bin_Code,
				Item_No:              value.Item_No,
				Variant_Code:         value.Variant_Code,
				Unit_of_Measure_Code: value.Unit_of_Measure_Code,
				Fixed:                value.Fixed,
				Default:              value.Default,
				Dedicated:            value.Dedicated,
				CalcQtyUOM:           value.CalcQtyUOM,
				Quantity_Base:        value.Quantity_Base,
				Bin_Type_Code:        value.Bin_Type_Code,
				Zone_Code:            value.Zone_Code,
				Lot_No_Filter:        value.Lot_No_Filter,
				Serial_No_Filter:     value.Serial_No_Filter,
				Branch:               &branchName,
			}

			allResults = append(allResults, maliwan_data)
		}
	}

	// Insert data in chunks
	chunkSize := 1000 // Adjust the chunk size as needed
	for i := 0; i < len(allResults); i += chunkSize {
		end := i + chunkSize
		if end > len(allResults) {
			end = len(allResults)
		}
		chunk := allResults[i:end]

		for j := range chunk {
			chunk[j].ItemabinmaliwanID = uuid.New().String()
		}

		db := database.DB
		tx := db.Create(&chunk)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	fmt.Println(len(allResults))
	return allResults, nil
}

func CreateNewMaliwan_data(maliwan_data model.Item_bin_maliwan) (value *model.Item_bin_maliwan, Error error) {
	db := database.DB
	maliwan_data.ItemabinmaliwanID = uuid.New().String()

	tx := db.Create(&maliwan_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_data, nil
}

func PullItembinHandler(c *fiber.Ctx) error {
	result, err := PullItembin(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(result)
}

// ---------------------------------------------------------------------------------------------------------------------------------

// func ReaddataAndCreateMaliwanbin(c *fiber.Ctx) error {
// 	values, err := PullItembin(c)
// 	if err != nil {
// 		return err
// 	}

// 	for _, value := range values {
// 		maliwan_data := model.Item_bin_maliwan{
// 			Location_Code:        value.Location_Code,
// 			Bin_Code:             value.Bin_Code,
// 			Item_No:              value.Item_No,
// 			Variant_Code:         value.Variant_Code,
// 			Unit_of_Measure_Code: value.Unit_of_Measure_Code,
// 			Fixed:                value.Fixed,
// 			Default:              value.Default,
// 			Dedicated:            value.Dedicated,
// 			CalcQtyUOM:           value.CalcQtyUOM,
// 			Quantity_Base:        value.Quantity_Base,
// 			Bin_Type_Code:        value.Bin_Type_Code,
// 			Zone_Code:            value.Zone_Code,
// 			Lot_No_Filter:        value.Lot_No_Filter,
// 			Serial_No_Filter:     value.Serial_No_Filter,
// 			Branch:               value.Branch,
// 		}

// 		_, err := CreateNewMaliwan_data(maliwan_data)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func ReaddataAndCreate(maliwanData []*model.Maliwan_data, maliwanCountID string) error {

// 	for _, value := range maliwanData {
// 		maliwan_data := model.Maliwan_Count_Store{
// 			Location_Code:              value.Location_Code,
// 			Bin_Code:                   value.Bin_Code,
// 			No_2:                       value.No_2,
// 			Description:                value.Description,
// 			Description_2:              value.Description_2,
// 			Description_ENG:            value.Description_ENG,
// 			Inventory:                  value.Inventory,
// 			Maximum_Inventory:          value.Safety_Stock_Quantity,
// 			Product_Group_Code:         value.Product_Group_Code,
// 			Created_From_Nonstock_Item: value.Created_From_Nonstock_Item,
// 			Line_Type:                  value.Line_Type,
// 			Type_Insurance:             value.Type_Insurance,
// 			Type_Registration_fee:      value.Type_Registration_fee,
// 			Type_Accessory:             value.Type_Accessory,
// 			Type_Deposit:               value.Type_Deposit,
// 			Stockkeeping_Unit_Exists:   value.Stockkeeping_Unit_Exists,
// 			Assembly_BOM:               value.Assembly_BOM,
// 			Base_Unit_of_Measure:       value.Base_Unit_of_Measure,
// 			Shelf_No:                   value.Shelf_No,
// 			Gen_Prod_Posting_Group:     value.Gen_Prod_Posting_Group,
// 			Inventory_Posting_Group:    value.Inventory_Posting_Group,
// 			Item_Category_Code:         value.Item_Category_Code,
// 			VAT_Prod_Posting_Group:     value.VAT_Prod_Posting_Group,
// 			Item_Disc_Group:            value.Item_Disc_Group,
// 			Vendor_No:                  value.Vendor_No,
// 			Vendor_Item_No:             value.Vendor_Item_No,
// 			Tariff_No:                  value.Tariff_No,
// 			Search_Description:         value.Search_Description,
// 			Overhead_Rate:              value.Overhead_Rate,
// 			Blocked:                    value.Blocked,
// 			Last_Date_Modified:         value.Last_Date_Modified,
// 			Sales_Unit_of_Measure:      value.Sales_Unit_of_Measure,
// 			Replenishment_System:       value.Replenishment_System,
// 			Purch_Unit_of_Measure:      value.Purch_Unit_of_Measure,
// 			Lead_Time_Calculation:      value.Lead_Time_Calculation,
// 			Manufacturing_Policy:       value.Manufacturing_Policy,
// 			Flushing_Method:            value.Flushing_Method,
// 			Assembly_Policy:            value.Assembly_Policy,
// 			Item_Tracking_Code:         value.Item_Tracking_Code,
// 			Promotion_Model_Code:       value.Promotion_Model_Code,
// 			Default_Location_Code:      value.Default_Location_Code,
// 			Qty_on_Purch_Order:         value.Qty_on_Purch_Order,
// 			Qty_on_Service_Order:       value.Qty_on_Service_Order,
// 			Qty_on_Sales_Order:         value.Qty_on_Sales_Order,
// 			Unit_Price_Price:           value.Unit_Price_Price,
// 			Brand:                      value.Brand,
// 			Costing_Method:             value.Costing_Method,
// 			Stockout_Warning:           value.Stockout_Warning,
// 			Global_Dimension_1_Filter:  value.Global_Dimension_1_Filter,
// 			Global_Dimension_2_Filter:  value.Global_Dimension_2_Filter,
// 			Location_Filter:            value.Location_Filter,
// 			Drop_Shipment_Filter:       value.Drop_Shipment_Filter,
// 			Variant_Filter:             value.Variant_Filter,
// 			Lot_No_Filter:              value.Lot_No_Filter,
// 			Serial_No_Filter:           value.Serial_No_Filter,
// 			Branch:                     value.Branch,
// 			Date_Filter:                value.Date_Filter,
// 			Maliwan_countID:            maliwanCountID,
// 		}

// 		_, err := ServicesMaliwan_Counting_Trigger.CreateNewMaliwan_Count_StoreData(maliwan_data)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func GetDatamaliwan(c *fiber.Ctx) error {
// 	Client := http.Client{
// 		Transport: &httpntlm.NtlmTransport{
// 			Domain:   viper.GetString("nav.domain"),
// 			User:     viper.GetString("nav.username"),
// 			Password: viper.GetString("nav.password"),
// 		},
// 	}
// 	branchMaliwanCounts := [...]string{"(Head Office)"}
// 	Gen_prod_posting_Group := "PARTS"
// 	itemCategoryCodes := [...]string{"OILS", "BRAKE"}
// 	var records model.Itemapiodata
// 	var result []model.Itemapivalue // เพิ่มตัวแปร result เพื่อเก็บผลลัพธ์

// 	for _, branchMaliwanCount := range branchMaliwanCounts {
// 		fmt.Println(branchMaliwanCount, "ค่าที่รับมา")
// 		top := 7000
// 		skip := 0
// 		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
// 		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
// 		ItemApi := "/ItemsAPI"

// 		// สร้างเงื่อนไขของ filter
// 		filter := fmt.Sprintf("Gen_Prod_Posting_Group eq '%v' and Inventory gt 0", Gen_prod_posting_Group)
// 		fmt.Println("Gen_prod_posting_Group = ", Gen_prod_posting_Group)
// 		var itemCategoryFilter string
// 		for i, code := range itemCategoryCodes {
// 			if i > 0 {
// 				itemCategoryFilter += " or "
// 			}
// 			itemCategoryFilter += fmt.Sprintf("Item_Category_Code eq '%v'", code)
// 		}
// 		if itemCategoryFilter != "" {
// 			filter += fmt.Sprintf(" and (%v)", itemCategoryFilter)
// 			fmt.Println("itemCategoryFilter = ", itemCategoryFilter)
// 		}

// 		param := fmt.Sprintf("?$top=%v&$skip=%v&$filter=%v", top, skip, url.QueryEscape(filter))
// 		path_url2 := NAV_SERVER_HMW + Company + ItemApi + param
// 		req, err := http.NewRequest("GET", path_url2, strings.NewReader(""))

// 		if err != nil {
// 			return err
// 		}

// 		for {
// 			resp, err := Client.Do(req)
// 			if err != nil {
// 				return err
// 			}
// 			defer resp.Body.Close()
// 			if resp.StatusCode != 200 {
// 				fmt.Println(resp.StatusCode)
// 				break
// 			}
// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				return err
// 			}
// 			err = json.Unmarshal(body, &records)
// 			if err != nil {
// 				return err
// 			}

// 			result = append(result, records.Value...)

// 			if len(records.Value) < top {
// 				break
// 			}
// 			skip += top

// 			param = fmt.Sprintf("?$top=%v&$skip=%v&$filter=%v", top, skip, url.QueryEscape(filter))
// 			path_url2 = NAV_SERVER_HMW + Company + ItemApi + param
// 			req, err = http.NewRequest("GET", path_url2, strings.NewReader(""))
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	fmt.Println(len(result))
// 	return c.JSON(result)
// }
