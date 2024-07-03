package ControllerCount_maliwan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ServicesMaliwan_Fixed_Asset_Store "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_store/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Fixed_Asset_Store.GetById(c.Params("maliwan_fixed_asset_store_id"))
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
	records, err := ServicesMaliwan_Fixed_Asset_Store.GetAllRecords()
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

func GetAllMaliwan_Fixed_AssetHandler(c *fiber.Ctx) error {
	records, err := ServicesMaliwan_Fixed_Asset_Store.GetAllMaliwan_Fixed_Asset()
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
	count_maliwan := new(model.Maliwan_Fixed_Asset_Store)
	err := c.BodyParser(count_maliwan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(count_maliwan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_Fixed_Asset_Store.CreateNewMaliwan_Fixed_Asset_Store(*count_maliwan)
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

func UpdateMaliwan_Fixed_Asset_Store(c *fiber.Ctx) error {
	id := c.Params("maliwan_fixed_asset_store_id")

	updatedCount_maliwan := new(model.Maliwan_Fixed_Asset_Store)
	err := c.BodyParser(updatedCount_maliwan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedCount_maliwan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := ServicesMaliwan_Fixed_Asset_Store.UpdateMaliwan_Fixed_Asset_Store(id, *updatedCount_maliwan)
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

func GetByMaliwan_Fixed_AssetC_CountIDHandler(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Fixed_Asset_Store.GetByMaliwan_Fixed_AssetC_CountID(c.Params("maliwan_fixed_asset_count_id"))
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

func DeleteMaliwan_Fixed_Asset_Store(c *fiber.Ctx) error {
	countingTriggerID := c.Params("maliwan_fixed_asset_count_id")

	err := ServicesMaliwan_Fixed_Asset_Store.DeleteMaliwan_Fixed_Asset_Store(countingTriggerID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "deleted successfully",
	})
}

func PullDataHandler(c *fiber.Ctx) error {
	if err := ServicesMaliwan_Fixed_Asset_Store.DeleteAllDataMaliwan(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error deleting all data",
			"error":   err.Error(),
		})
	}

	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	branchMaliwanCounts := []string{"(Head Office)", "(Mueang Krabi)", "(Naka)", "(Burirum)", "(Surin)"}

	for _, branchMaliwanCount := range branchMaliwanCounts {
		fmt.Println(branchMaliwanCount, "ค่าที่รับมา")
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/FixedAssetListAPI"

		var result []model.Maliwan_Fixed_Asset

		for {
			param := fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
			pathURL2 := NAV_SERVER_HMW + Company + ItemApi + param
			req, err := http.NewRequest("GET", pathURL2, nil)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error creating request",
					"error":   err.Error(),
				})
			}

			resp, err := Client.Do(req)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error making request",
					"error":   err.Error(),
				})
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
				break
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error reading response body",
					"error":   err.Error(),
				})
			}

			var records model.Item_Maliwan_Fixed_Asset
			err = json.Unmarshal(body, &records)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error unmarshalling response",
					"error":   err.Error(),
				})
			}

			result = append(result, records.Value...)

			if len(records.Value) < top {
				break
			}
			skip += top
		}

		// Insert data for this branch in chunks
		chunkSize := 1000 // Adjust the chunk size as needed
		for i := 0; i < len(result); i += chunkSize {
			end := i + chunkSize
			if end > len(result) {
				end = len(result)
			}
			chunk := result[i:end]

			// Create a new UUID for each chunk
			for j := range chunk {
				chunk[j].Maliwan_Fixed_AssetId = uuid.New().String()
				chunk[j].Branch = strings.Trim(branchMaliwanCount, "()")
			}

			// Insert the chunk into the database
			db := database.DB
			tx := db.Create(&chunk)
			if tx.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error creating new fixed asset store",
					"error":   tx.Error.Error(),
				})
			}
		}
	}

	_, err := ServicesMaliwan_Fixed_Asset_Store.UpdateRequest_Update_Maliwan_Fixed_Asset()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error updating request update data",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}

// func fetchFixedAssets(locationCode string) (result []model.ItemMaliwan_Fixed_Asset, err error) {
// 	Client := http.Client{
// 		Transport: &httpntlm.NtlmTransport{
// 			User:     viper.GetString("nav.username"),
// 			Password: viper.GetString("nav.password"),
// 		},
// 	}

// 	NAV_SERVER_MALIWAN := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Maliwan')"
// 	ItemAPI := "/FixedAssetListAPI"
// 	filter := fmt.Sprintf("?$filter=FA_Location_Code%%20eq%%20'%s'", locationCode)

// 	path_url := NAV_SERVER_MALIWAN + ItemAPI + filter
// 	req, err := http.NewRequest("GET", path_url, strings.NewReader(""))
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := Client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var maliwan []model.ItemMaliwan_Fixed_Asset
// 	err = json.Unmarshal(body, &maliwan)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return maliwan, nil
// }

func Import_data_Store(c *fiber.Ctx) error {
	value, err := ServicesMaliwan_Fixed_Asset_Store.GetByMaliwan_Fixed_Asset_Count(c.Params("maliwan_fixed_asset_count_id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Data not found",
			"error":   err.Error(),
		})
	}

	maliwan_fixed_asset_count_id := value.Maliwan_Fixed_Asset_CountID

	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	branchMaliwanCounts := strings.Split(value.Branch, ",")
	for i, branch := range branchMaliwanCounts {
		branchMaliwanCounts[i] = fmt.Sprintf("(%s)", strings.TrimSpace(branch))
	}
	var result []model.Maliwan_Fixed_Asset

	for _, branchMaliwanCount := range branchMaliwanCounts {
		result = nil

		fmt.Println(branchMaliwanCount, "ค่าที่รับมา")
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/FixedAssetListAPI"

		param := fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
		pathURL2 := NAV_SERVER_HMW + Company + ItemApi + param
		req, err := http.NewRequest("GET", pathURL2, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Error creating request",
				"error":   err.Error(),
			})
		}

		for {
			resp, err := Client.Do(req)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error making request",
					"error":   err.Error(),
				})
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
				break
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error reading response body",
					"error":   err.Error(),
				})
			}

			var records model.Item_Maliwan_Fixed_Asset
			err = json.Unmarshal(body, &records)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error unmarshalling response",
					"error":   err.Error(),
				})
			}

			result = append(result, records.Value...)

			if len(records.Value) < top {
				break
			}
			skip += top

			param = fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
			pathURL2 = NAV_SERVER_HMW + Company + ItemApi + param
			req, err = http.NewRequest("GET", pathURL2, nil)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error creating request",
					"error":   err.Error(),
				})
			}
		}

		for _, value := range result {

			branchName := strings.Trim(branchMaliwanCount, "()")

			maliwanData := model.Maliwan_Fixed_Asset_Store{
				No:                          value.No,
				Description:                 value.Description,
				Description_2:               value.Description_2,
				Ref_Code:                    value.Ref_Code,
				Serial_No:                   value.Serial_No,
				Search_Description:          value.Search_Description,
				Responsible_Employee:        value.Responsible_Employee,
				FA_Posting_Group:            value.FA_Posting_Group,
				FA_Class_Code:               value.FA_Class_Code,
				FA_Subclass_Code:            value.FA_Subclass_Code,
				FA_Location_Code:            value.FA_Location_Code,
				FA_Department_Code:          value.FA_Department_Code,
				Budgeted_Asset:              value.Budgeted_Asset,
				Inactive:                    value.Inactive,
				Blocked:                     value.Blocked,
				Global_Dimension_1_Code:     value.Global_Dimension_1_Code,
				Global_Dimension_2_Code:     value.Global_Dimension_2_Code,
				Main_Asset_Component:        value.Main_Asset_Component,
				Component_of_Main_Asset:     value.Component_of_Main_Asset,
				Branch:                      branchName,
				Maliwan_Fixed_Asset_CountID: maliwan_fixed_asset_count_id,
			}

			// fmt.Println(branchName, "สาขาที่รับค่ามาใน store")

			_, err := ServicesMaliwan_Fixed_Asset_Store.CreateNewMaliwan_Fixed_Asset_Store(maliwanData)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Error creating new fixed asset store",
					"error":   err.Error(),
				})
			}
		}
	}

	_, err = ServicesMaliwan_Fixed_Asset_Store.UpdateMaliwanCountStatusPullData(maliwan_fixed_asset_count_id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetByNoHandler(c *fiber.Ctx) error {
	// Extract the query parameter 'no' directly from the URL query string
	no := c.Query("no")

	// Check if 'no' parameter is empty
	if no == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "No 'no' parameter provided",
		})
	}

	// Call the GetByNoDB function passing the 'no' parameter as argument
	value, err := ServicesMaliwan_Fixed_Asset_Store.GetByNoDB(no)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetByBranchHandler(c *fiber.Ctx) error {
	// Extract the query parameter 'no' directly from the URL query string
	branch := c.Query("branch")

	// Check if 'no' parameter is empty
	if branch == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Branch 'branch' parameter provided",
		})
	}

	// Call the GetByNoDB function passing the 'no' parameter as argument
	value, err := ServicesMaliwan_Fixed_Asset_Store.GetByBranchDB(branch)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func CreateMaliwan_Update_Fixed_Asset_StoryHandler(c *fiber.Ctx) error {
	count_maliwan := new(model.Maliwan_Update_Fixed_Asset_Story)
	err := c.BodyParser(count_maliwan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(count_maliwan)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMaliwan_Fixed_Asset_Store.CreateMaliwan_Update_Fixed_Asset_Story(*count_maliwan)
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

func GetAllMaliwan_Update_Fixed_Asset_StoryHandler(c *fiber.Ctx) error {
	records, err := ServicesMaliwan_Fixed_Asset_Store.GetAllMaliwan_Update_Fixed_Asset_Story()
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

func GetMaliwan_Fixed_AssetsHandler(c *fiber.Ctx) error {
	// อ่านพารามิเตอร์ limit จาก URL query
	limit, err := strconv.Atoi(c.Query("limit", "1000"))
	if err != nil || limit < 1 {
		limit = 1000 // กำหนดค่าเริ่มต้นให้เป็น 1000 ถ้า limit ไม่ถูกต้อง
	}

	// อ่านพารามิเตอร์ page จาก URL query
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1 // กำหนดค่าเริ่มต้นให้เป็น 1 ถ้า page ไม่ถูกต้อง
	}
	offset := (page - 1) * limit

	// ดึงข้อมูลจากฐานข้อมูลโดยใช้ limit และ offset ที่กำหนด
	records, err := ServicesMaliwan_Fixed_Asset_Store.GetMaliwan_Fixed_AssetDB(limit, offset)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	// ดึงจำนวนข้อมูลทั้งหมดจากฐานข้อมูล
	totalRecords, err := ServicesMaliwan_Fixed_Asset_Store.GetTotalMaliwan_Fixed_Assets()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve total count!",
			"error":   err,
		})
	}

	// ส่งผลลัพธ์กลับในรูปแบบ JSON
	return c.JSON(fiber.Map{
		"status":      "success",
		"result":      records,
		"page":        page,
		"limit":       limit,
		"total_count": totalRecords,
	})
}
