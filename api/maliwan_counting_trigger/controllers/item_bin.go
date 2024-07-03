package ControllerMaliwan_Counting_Trigger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	ServicesMaliwan_Counting_Trigger "github.com/puvadon-artmit/gofiber-template/api/maliwan_counting_trigger/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func DeleteItem_bin_maliwan(c *fiber.Ctx) error {

	if err := ServicesMaliwan_Counting_Trigger.DeleteAllItem_bin_maliwan(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All records in Item_bin_maliwan have been deleted",
	})
}

func FilterItembin(branchMaliwanCounts []string, itemNos []string) ([]model.Item_bin_maliwan, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	var result []model.Item_bin_maliwan

	for _, branchMaliwanCount := range branchMaliwanCounts {
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/ItemBinCodeAPI"

		filter := ""
		for i, itemNo := range itemNos {
			if i > 0 {
				filter += " or "
			}
			filter += fmt.Sprintf("Item_No eq '%s'", url.PathEscape(itemNo))
		}

		param := fmt.Sprintf("?$top=%v&$skip=%v&$filter=(%s)", top, skip, filter)
		path_url := NAV_SERVER_HMW + Company + ItemApi + param

		fmt.Println(path_url)

		req, err := http.NewRequest("GET", path_url, nil)
		if err != nil {
			return nil, err
		}

		for {
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

			param = fmt.Sprintf("?$top=%v&$skip=%v&$filter=(%s)", top, skip, filter)
			path_url = NAV_SERVER_HMW + Company + ItemApi + param
			req, err = http.NewRequest("GET", path_url, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	fmt.Println(len(result))
	return result, nil
}

// ---------------------------------------------------------------------------------------------------------------------------

func GetByMaliwan_CountIDOnFilterODataV4(c *fiber.Ctx) error {
	maliwanCountID := c.Params("maliwan_count_id")

	if err := FilterDataODataV4(c, maliwanCountID); err != nil {
		return err
	}

	return nil

	// รอให้ PullDataByBranch ฟังชั่นนี้ทำงานเสร็จก่อนทำงาน ProcessmaliwanCountID
	// return ProcessmaliwanCountID(c, maliwanCountID)
}

func FilterDataODataV4(c *fiber.Ctx, maliwanCountID string) error {
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

	// สร้าง slice เพื่อเก็บข้อมูล branchMaliwanCount เป็นสมาชิกเดียว
	ArrayofCompanyName := []string{fmt.Sprintf("(%s)", *branchMaliwanCount)}

	response := fmt.Sprintf("(%s)", *branchMaliwanCount)

	err = GetAutoByTrigger_IDODataV4(c, ArrayofCompanyName, Gen_prod_posting_Group_Maliwan, itemCategoryCodeList, maliwanCountID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	fmt.Println(response)
	return nil
}

func GetAutoByTrigger_IDODataV4(c *fiber.Ctx, branchMaliwanCounts []string, Gen_prod_posting_Group string, itemCategoryCodes []string, maliwanCountID string) error {
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
			branchName := branchMaliwanCount

			var NoSlice []string
			for _, value := range result {
				NoSlice = append(NoSlice, value.No)
			}

			// เรียกใช้ฟังก์ชัน GetAllbranchsMaliwan โดยส่งข้อมูลที่จำเป็น
			maliwan, err := FilterItembin([]string{branchName}, NoSlice)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}

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

				_, err := ServicesMaliwan_Counting_Trigger.CreateMaliwan_Count_Store(maliwan_data)
				if err != nil {
					return err
				}
			}

			fmt.Println(len(maliwan))
			return c.JSON(maliwan)
		}
	}
	return nil
}
