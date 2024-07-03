package ItemAutoclik_Fixed_AssetServices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/api/redis"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
	"gorm.io/gorm/clause"
)

func GetAutoclik_Fixed_Asset() (result []model.ItemAutoclik_Fixed_Asset, err error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/FixedAssetListAPI"
	path_url := NAV_SERVER_AUTOCLIK + ItemAPI
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

	var autoclik model.ItemAutoclik_Fixed_Asset
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return nil, err
	}

	return []model.ItemAutoclik_Fixed_Asset{autoclik}, nil
}

func CreateAutoclik_Fixed_Asset(autoclik_data model.Autoclik_Fixed_Asset) (value *model.Autoclik_Fixed_Asset, Error error) {
	db := database.DB
	autoclik_data.Autoclik_Fixed_AssetId = uuid.New().String()
	tx := db.Create(&autoclik_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_data, nil
}

func GetAllAutoclik_Fixed_Asset() (result []*model.Autoclik_Fixed_Asset, err error) {
	cachedData, err := redis.GetAutoclik_Fixed_Assetcache("autoclik_fixed_asset")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		return cachedData, nil
	}

	db := database.DB
	var records []*model.Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(records) == 0 {
		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
		return records, nil
	}

	err = redis.SetAutoclik_Fixed_Assetcache("autoclik_fixed_asset", records)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}

	return records, nil
}

func CountLatestAutoclik_data() (count int, err error) {

	cachedData, err := redis.GetAutoclik_Fixed_Assetcache("autoclik_fixed_asset")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		return len(cachedData), nil
	}

	latestAutoclik_data, err := GetAllAutoclik_Fixed_Asset()
	if err != nil {
		return 0, err
	}
	err = redis.SetAutoclik_Fixed_Assetcache("autoclik_fixed_asset", latestAutoclik_data)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}
	return len(latestAutoclik_data), nil
}

func GetFA_Location_Code() (faLocationCodes []string, faPostingGroups []string, err error) {
	db := database.DB
	var records []*model.Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}

	seenLocation := make(map[string]bool)
	seenGroup := make(map[string]bool)

	for _, record := range records {
		if record.FA_Location_Code != nil && *record.FA_Location_Code != "" {
			if _, ok := seenLocation[*record.FA_Location_Code]; !ok {
				faLocationCodes = append(faLocationCodes, *record.FA_Location_Code)
				seenLocation[*record.FA_Location_Code] = true
			}
		}
		if record.FA_Posting_Group != nil && *record.FA_Posting_Group != "" {
			if _, ok := seenGroup[*record.FA_Posting_Group]; !ok {
				faPostingGroups = append(faPostingGroups, *record.FA_Posting_Group)
				seenGroup[*record.FA_Posting_Group] = true
			}
		}
	}

	sort.Strings(faLocationCodes)
	sort.Strings(faPostingGroups)

	return faLocationCodes, faPostingGroups, nil
}

func GetByItem_Autoclik_BinNoDB(No string) ([]*model.Autoclik_Fixed_Asset, error) {
	db := database.DB
	var autoclik []*model.Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Where("no = ?", No).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik, nil
}

func GetByProductGroupNoDB(No string) ([]*model.Item_Autoclik_Bin_Code, error) {
	db := database.DB
	var product_group []*model.Item_Autoclik_Bin_Code
	tx := db.Preload(clause.Associations).Where("product_group = ?", No).Find(&product_group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product_group, nil
}

func CountAllItemBin() (count int, err error) {
	db := database.DB
	var totalCount int64
	tx := db.Model(&model.Item_Autoclik_Bin_Code{}).Count(&totalCount)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(totalCount), nil
}

func GetAllRecordsAndCreateAutoclik_Fixed_Asset() error {
	values, err := GetAutoclik_Fixed_Asset()
	if err != nil {
		return err
	}

	var autoclikData []model.Autoclik_Fixed_Asset
	for _, value := range values {
		for _, item := range value.Value {
			autoclikData = append(autoclikData, model.Autoclik_Fixed_Asset{
				Autoclik_Fixed_AssetId:  uuid.New().String(), // Generate UUID for each record
				No:                      item.No,
				Description:             item.Description,
				Description_2:           item.Description_2,
				Ref_Code:                item.Ref_Code,
				Serial_No:               item.Serial_No,
				Search_Description:      item.Search_Description,
				Responsible_Employee:    item.Responsible_Employee,
				FA_Posting_Group:        item.FA_Posting_Group,
				FA_Class_Code:           item.FA_Class_Code,
				FA_Subclass_Code:        item.FA_Subclass_Code,
				FA_Location_Code:        item.FA_Location_Code,
				FA_Department_Code:      item.FA_Department_Code,
				Budgeted_Asset:          item.Budgeted_Asset,
				Inactive:                item.Inactive,
				Blocked:                 item.Blocked,
				Global_Dimension_1_Code: item.Global_Dimension_1_Code,
				Global_Dimension_2_Code: item.Global_Dimension_2_Code,
				Main_Asset_Component:    item.Main_Asset_Component,
				Component_of_Main_Asset: item.Component_of_Main_Asset,
			})
		}
	}

	// Insert data in chunks
	chunkSize := 1000 // Adjust the chunk size as needed
	for i := 0; i < len(autoclikData); i += chunkSize {
		end := i + chunkSize
		if end > len(autoclikData) {
			end = len(autoclikData)
		}
		chunk := autoclikData[i:end]

		db := database.DB
		tx := db.Create(&chunk)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func ClearAllAutoclik_Fixed_Asset() error {
	db := database.DB
	if err := db.Where("true").Delete(&model.Autoclik_Fixed_Asset{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequest_Update_Autoclik_Fixed_Asset() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "9a11b454-7d3f-405d-95a4-9443e6835d77"
	existingType := model.Request_Update_Data{Request_Update_DataID: id}

	tx := db.Preload(clause.Associations).First(&existingType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	succeed := "succeed"
	tx = db.Model(&existingType).Update("Status", succeed)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func ClearAndFetchAllAutoclik_Fixed_Asset() error {
	// เคลียร์ข้อมูลทั้งหมดจากตาราง
	if err := ClearAllAutoclik_Fixed_Asset(); err != nil {
		return err
	}

	if err := redis.DeleteKey("autoclik_fixed_asset"); err != nil {
		return err
	}

	// สร้างข้อมูลใหม่จากข้อมูลที่มีอยู่
	if err := GetAllRecordsAndCreateAutoclik_Fixed_Asset(); err != nil {
		return err
	}

	_, err := UpdateRequest_Update_Autoclik_Fixed_Asset()
	if err != nil {
		return err
	}

	return nil
}

func Filter_Data(FA_Posting_Group []string, FA_Location_Code string) ([]*model.Autoclik_Fixed_Asset, error) {
	db := database.DB
	var product_group []*model.Autoclik_Fixed_Asset
	tx := db.Preload(clause.Associations).Where("fa_posting_group IN (?) AND fa_location_code = ?", FA_Posting_Group, FA_Location_Code).Find(&product_group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product_group, nil
}

// func GetChoosebytype() (ItemCategory []string, ProductGroup []string, Location_Code []string, err error) {
// 	db := database.DB
// 	var autoclik []*model.Item_Autoclik_Bin_Code
// 	tx := db.Select("ItemCategory", "ProductGroup", "Location_Code").Preload(clause.Associations).Find(&autoclik)
// 	if tx.Error != nil {
// 		return nil, nil, nil, tx.Error
// 	}

// 	ItemCategoryMap := make(map[string]struct{})
// 	ProductGroupMap := make(map[string]struct{})
// 	Location_CodeMap := make(map[string]struct{})
// 	for _, data := range autoclik {
// 		addToMapIfNotNil(ItemCategoryMap, data.ItemCategory)
// 		addToMapIfNotNil(ProductGroupMap, data.ProductGroup)
// 		addToMapIfNotNil(Location_CodeMap, data.Location_Code)
// 	}

// 	ItemCategory = mapToSortedSlice(ItemCategoryMap)
// 	ProductGroup = mapToSortedSlice(ProductGroupMap)
// 	Location_Code = mapToSortedSlice(Location_CodeMap)

// 	return ItemCategory, ProductGroup, Location_Code, nil
// }

func GetAutoclik_Fixed_AssetFilter(locationCode string) (result []model.ItemAutoclik_Fixed_Asset, err error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/FixedAssetListAPI"
	filter := fmt.Sprintf("?&$filter=FA_Location_Code%%20eq%%20'%s'%%20and%%20(", locationCode)

	path_url := NAV_SERVER_AUTOCLIK + ItemAPI + filter
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

	var autoclik model.ItemAutoclik_Fixed_Asset
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return nil, err
	}

	return []model.ItemAutoclik_Fixed_Asset{autoclik}, nil
}
