package ControllerAutoclik_Counting_Trigger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Counting_Trigger "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_trigger/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

// func Masterdata() (result []model.Itemapivalue, err error) {

// 	Client := http.Client{
// 		Transport: &httpntlm.NtlmTransport{
// 			Domain:   viper.GetString("nav.domain"),
// 			User:     viper.GetString("nav.username"),
// 			Password: viper.GetString("nav.password"),
// 		},
// 	}

// 	var autoclik model.Itemapiodata
// 	top := 1000
// 	skip := 0
// 	NAV_SERVER_AUTOCLIK := viper.GetString("nav.odata_autoclik_server")
// 	ItemAPI := viper.GetString("nav.item_api")
// 	param := fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
// 	path_url := NAV_SERVER_AUTOCLIK + ItemAPI + param
// 	req, err := http.NewRequest("GET", path_url, strings.NewReader(""))
// 	for {
// 		resp, err := Client.Do(req)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer resp.Body.Close()
// 		if resp.StatusCode != 200 {
// 			fmt.Println(resp.StatusCode)
// 			break
// 		}
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		err = json.Unmarshal(body, &autoclik)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, autoclik.Value...)
// 		if len(autoclik.Value) < top {
// 			break
// 		}
// 		skip += top
// 		NAV_SERVER_AUTOCLIK := viper.GetString("nav.odata_autoclik_server")
// 		ItemAPI := viper.GetString("nav.item_api")
// 		param := fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
// 		path_url := NAV_SERVER_AUTOCLIK + ItemAPI + param
// 		req, err = http.NewRequest("GET", path_url, strings.NewReader(""))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	fmt.Println(len(result))
// 	return result, nil
// }

func Itembin(locationCode string, productGroups []string) (result []model.Itemapiobindata, err error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/ItemBinCodeAPI"

	var filters []string
	for _, pg := range productGroups {
		filters = append(filters, fmt.Sprintf("ProductGroup eq '%s'", pg))
	}
	filterStr := fmt.Sprintf("Location_Code eq '%s' and (%s)", locationCode, strings.Join(filters, " or "))

	path_url := fmt.Sprintf("%s%s?filter=%s", NAV_SERVER_AUTOCLIK, ItemAPI, filterStr)

	req, err := http.NewRequest("GET", path_url, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var autoclik model.Itemapiobindata
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return nil, err
	}

	return []model.Itemapiobindata{autoclik}, nil
}

func GetItem_binJSonData(c *fiber.Ctx) error {
	locationCode := c.Query("location_code")
	productGroups := c.Query("product_group")
	productGroupList := strings.Split(productGroups, ",")

	masterData, err := Itembin(locationCode, productGroupList)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"Message": "Failed to fetch master data!",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": masterData,
	})
}

func GetByAutoclik_Count_ID2(c *fiber.Ctx) error {
	Autoclik_Count_ID := c.Params("autoclik_count_id")

	return ProcessAutoclikCountID(c, Autoclik_Count_ID)
}

func ProcessAutoclikCountID2(c *fiber.Ctx, Autoclik_Count_ID string) error {
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

		// นำ Gen_Prod_Posting_Group ออกมาจากใบใหญ่เพื่อนำไป filter
		genProdPostingGroup := autoclikCountingTrigger.Gen_Prod_Posting_Group

		// Ensure branchAutoclik has at least one element
		if len(branchAutoclik) == 0 {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch data not found",
			})
		}

		// นำ Branch_Code ออกมาจากใบใหญ่เพื่อนำไป filter
		branchCode := branchAutoclik[0].Branch_Autoclik.Branch_Code
		if branchCode == "" {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  "Branch code is empty",
			})
		}

		// นำ Name_Product_Group ออกมาจากใบใหญ่เพื่อนำไป filter
		var nameProductGroups []string
		for _, pg := range product_groupAutoclik {
			if pg.Name_Product_Group != nil {
				nameProductGroups = append(nameProductGroups, *pg.Name_Product_Group)
			}
		}

		// filter gen_prod_posting_group ใน autoclik_data
		genProdPostingGroupResults, err := ServicesAutoclik_Counting_Trigger.GetByGen_Prod_Posting_GroupDB(genProdPostingGroup)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		//filter product_group และ location_code/Location_Code ใน autoclik_bin
		productGroups, err := ServicesAutoclik_Counting_Trigger.GetByProductGroupNoDB(nameProductGroups, branchCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Filter Product_Groups based on the condition Item_No == No เพื่อออกไปแสดงผล
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
				"product_group":          filteredProductGroups,
			},
		}

		return c.JSON(response)
	}

	return c.Status(404).JSON(fiber.Map{
		"status": "error",
		"error":  "Data not found",
	})
}
