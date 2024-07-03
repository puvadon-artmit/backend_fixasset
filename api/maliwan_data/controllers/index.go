package Maliwan_dataController

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
	Maliwan_datavices "github.com/puvadon-artmit/gofiber-template/api/maliwan_data/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func GetAllGetItem_Category_Code(c *fiber.Ctx) error {
	ctx := c.Context() // สร้าง context จาก fiber.Ctx

	Item_Category_Code, Branch, err := Maliwan_datavices.GetAllMaliwan_Item_Category_Code_And_Filter(ctx)
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
			"Item_Category_Code": Item_Category_Code,
			"Branch":             Branch,
		},
	})
}

func Clearandsuckbmaliwan(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndPullMaliwan_Item()
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

func GetAllDataMaliwanHandler(c *fiber.Ctx) error {
	records, err := Maliwan_datavices.GetAllMaliwanData()
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

func GetCountAllMaliwan(c *fiber.Ctx) error {
	count, err := Maliwan_datavices.CountAllMaliwan()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("%d", count))
}

func GetAllHandler(c *fiber.Ctx) error {
	Branch := c.Query("branch")

	fmt.Println(string(Branch))
	value, err := Maliwan_datavices.GetAllMaliwan(Branch)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(value)
}

func GetCountMaliwandata(c *fiber.Ctx) error {
	Branch := c.Query("branch")
	count, err := Maliwan_datavices.GetByCountBranch(Branch)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)
}

func GetByCountBranchHandler(c *fiber.Ctx) error {
	Branch := c.Query("branch")

	fmt.Println(string(Branch))
	value, err := Maliwan_datavices.GetByCountBranch(Branch)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(value)
}

func GetByNoHandler(c *fiber.Ctx) error {
	No := c.Query("no")
	maliwan, err := Maliwan_datavices.GetByNoDB(No)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": maliwan,
	})
}

func GetByBranch(c *fiber.Ctx) error {
	Branch := c.Query("branch")
	Gen_prod_posting_group := c.Query("gen_prod_posting_group")

	fmt.Println(string(Branch))
	value, err := Maliwan_datavices.GetAllbranchsMaliwan(Branch, Gen_prod_posting_group)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(value)
}

func GetByBranchHandler(c *fiber.Ctx) error {
	Branch := c.Query("branch")
	Limit, _ := strconv.Atoi(c.Query("limit"))
	Offset, _ := strconv.Atoi(c.Query("offset"))

	fmt.Println(string(Branch))
	value, err := Maliwan_datavices.GetByBranch(Branch, Limit, Offset)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(value)
}

func GetCountMaliwanDataByBranch(c *fiber.Ctx) error {
	records, err := Maliwan_datavices.CountMaliwanDataByBranch()
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

// -----------------------------------------------------------------------------------------------------------------------------

func ClearandpullHeadOffice(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndFetchHeadOfficeItemBin()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data HeadOffice retrieved and created successfully",
	})
}

// ------------------------------------------------------------------------------------------------------------------------------

func ClearandpullBurirum(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndFetchBurirumItemBin()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Burirum  retrieved and created successfully",
	})
}

// ------------------------------------------------------------------------------------------------------------------------------

func ClearandpullNaka(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndFetchNakaItemBin()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Naka  retrieved and created successfully",
	})
}

// ------------------------------------------------------------------------------------------------------------------------------

func ClearandpullMueang_Krabi(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndFetchMueang_KrabiItemBin()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Mueang Krabi  retrieved and created successfully",
	})
}

// ------------------------------------------------------------------------------------------------------------------------------

func ClearandpullSurin(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndFetchSurinItemBin()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Surin retrieved and created successfully",
	})
}

// ------------------------------------------------------------------------------------------------------------------------------

func GetByItem_Category_CodeHandler(c *fiber.Ctx) error {
	item_category_code := c.Query("Item_Category_Code")
	maliwan, err := Maliwan_datavices.GetByItem_Category_CodeDB(item_category_code)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": maliwan,
	})
}

func GetFilterRecords(c *fiber.Ctx) error {
	err := Maliwan_datavices.GetAllRecordsAndCreateItem_Category_Code()
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

func GetAllMaliwan_Item_Category_CodeHandler(c *fiber.Ctx) error {
	records, err := Maliwan_datavices.GetAllMaliwan_Item_Category_Code()
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

// ---------------------------------------------------------------------------------------------------------

func ClearandpullAllMaliwan_data(c *fiber.Ctx) error {
	err := Maliwan_datavices.ClearAndPullMaliwan_Item()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Data Surin retrieved and created successfully",
	})
}

func CreateMaliwan_Update_Story(c *fiber.Ctx) error {
	inspection_story := new(model.Maliwan_Update_Story)
	err := c.BodyParser(inspection_story)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(inspection_story)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := Maliwan_datavices.CreateNewMaliwan_Update_Story(*inspection_story)
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
	records, err := Maliwan_datavices.GetMaliwanDataDB(limit, offset)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	// ดึงจำนวนข้อมูลทั้งหมดจากฐานข้อมูล
	totalRecords, err := Maliwan_datavices.GetTotalMaliwanData()
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

func GetByBranchtestHandler(c *fiber.Ctx) error {
	branch := c.Query("branch")
	genProdPostingGroup := c.Query("gen_prod_posting_group")
	page := c.Query("page", "1")
	limit := 7000

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid page number",
		})
	}

	offset := (pageNumber - 1) * limit

	value, err := Maliwan_datavices.GetBatchOfBranchsMaliwan(branch, genProdPostingGroup, offset, limit)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(value)
}

func PullDataMaliwanAllBranchHandler(c *fiber.Ctx) error {
	if err := Maliwan_datavices.DeleteAllDataMaliwan(); err != nil {
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
		ItemApi := "/ItemsAPI"

		var result []model.Maliwan_data

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

			var records model.ItemaMaliwan_data
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
				chunk[j].Maliwan_dataID = uuid.New().String()
				branchName := strings.Trim(branchMaliwanCount, "()")
				chunk[j].Branch = &branchName
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

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
