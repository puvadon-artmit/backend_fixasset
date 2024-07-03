package ControllerCount_autoclik

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAutoclik_Fixed_Asset_Count "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_count/services"
	ServicesAutoclik_Fixed_Asset_Store "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_store/services"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Fixed_Asset_Store.GetById(c.Params("autoclik_fixed_asset_store_id"))
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
	records, err := ServicesAutoclik_Fixed_Asset_Store.GetAllRecords()
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
	count_autoclik := new(model.Autoclik_Fixed_Asset_Store)
	err := c.BodyParser(count_autoclik)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(count_autoclik)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAutoclik_Fixed_Asset_Store.CreateNewAutoclik_Fixed_Asset_Store(*count_autoclik)
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

func UpdateAutoclik_Fixed_Asset_Store(c *fiber.Ctx) error {
	id := c.Params("autoclik_fixed_asset_store_id")

	updatedCount_autoclik := new(model.Autoclik_Fixed_Asset_Store)
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

	updatedStatus, err := ServicesAutoclik_Fixed_Asset_Store.UpdateAutoclik_Fixed_Asset_Store(id, *updatedCount_autoclik)
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

func GetByAutoclik_Fixed_AssetC_CountIDHandler(c *fiber.Ctx) error {
	value, err := ServicesAutoclik_Fixed_Asset_Store.GetByAutoclik_Fixed_AssetC_CountID(c.Params("autoclik_fixed_asset_count_id"))
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

func DeleteAutoclik_Fixed_Asset_Store(c *fiber.Ctx) error {
	countingTriggerID := c.Params("autoclik_fixed_asset_count_id")

	err := ServicesAutoclik_Fixed_Asset_Store.DeleteAutoclik_Fixed_Asset_Store(countingTriggerID)
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

func GetByIdHandler(c *fiber.Ctx) error {
	data_filter, err := ServicesAutoclik_Fixed_Asset_Count.GetById(c.Params("autoclik_fixed_asset_count_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/FixedAssetListAPI"
	filter := fmt.Sprintf("?$filter=FA_Location_Code%%20eq%%20'%s'", *data_filter.FA_Location_Code)

	path_url := NAV_SERVER_AUTOCLIK + ItemAPI + filter
	req, err := http.NewRequest("GET", path_url, strings.NewReader(""))
	if err != nil {
		return err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var autoclik model.ItemAutoclik_Fixed_Asset
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return err
	}

	for _, autoclikItem := range autoclik.Value {

		autoclik := model.Autoclik_Fixed_Asset_Store{
			No:                           autoclikItem.No,
			Description:                  autoclikItem.Description,
			Description_2:                autoclikItem.Description_2,
			Ref_Code:                     autoclikItem.Ref_Code,
			Serial_No:                    autoclikItem.Serial_No,
			Search_Description:           autoclikItem.Search_Description,
			Responsible_Employee:         autoclikItem.Responsible_Employee,
			FA_Posting_Group:             autoclikItem.FA_Posting_Group,
			FA_Class_Code:                autoclikItem.FA_Class_Code,
			FA_Subclass_Code:             autoclikItem.FA_Subclass_Code,
			FA_Location_Code:             autoclikItem.FA_Location_Code,
			FA_Department_Code:           autoclikItem.FA_Department_Code,
			Budgeted_Asset:               autoclikItem.Budgeted_Asset,
			Inactive:                     autoclikItem.Inactive,
			Blocked:                      autoclikItem.Blocked,
			Global_Dimension_1_Code:      autoclikItem.Global_Dimension_1_Code,
			Global_Dimension_2_Code:      autoclikItem.Global_Dimension_2_Code,
			Main_Asset_Component:         autoclikItem.Main_Asset_Component,
			Component_of_Main_Asset:      autoclikItem.Component_of_Main_Asset,
			Autoclik_Fixed_Asset_CountID: &data_filter.Autoclik_Fixed_Asset_CountID,
		}

		_, err := ServicesAutoclik_Fixed_Asset_Count.CreateNewAutoclik_Fixed_Asset_Count_Store(autoclik)
		if err != nil {
			return err
		}
	}

	_, err = ServicesAutoclik_Fixed_Asset_Count.UpdateAutoclik_Fixed_AssetCountStatusPullData(data_filter.Autoclik_Fixed_Asset_CountID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	fmt.Println("ทำงานครบถ้วน")

	return c.JSON(fiber.Map{
		"status": "success",
		"result": fiber.Map{
			"autoclik_fixed_asset_count_id": data_filter.Autoclik_Fixed_Asset_CountID,
			"fa_location_code":              data_filter.FA_Location_Code,
			"assets":                        autoclik.Value,
		},
	})
}

// func fetchFixedAssets(locationCode string) (result []model.ItemAutoclik_Fixed_Asset, err error) {
// 	Client := http.Client{
// 		Transport: &httpntlm.NtlmTransport{
// 			User:     viper.GetString("nav.username"),
// 			Password: viper.GetString("nav.password"),
// 		},
// 	}

// 	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
// 	ItemAPI := "/FixedAssetListAPI"
// 	filter := fmt.Sprintf("?$filter=FA_Location_Code%%20eq%%20'%s'", locationCode)

// 	path_url := NAV_SERVER_AUTOCLIK + ItemAPI + filter
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

// 	var autoclik []model.ItemAutoclik_Fixed_Asset
// 	err = json.Unmarshal(body, &autoclik)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return autoclik, nil
// }
