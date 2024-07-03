package Autoclik_dataServices

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

func GetAllRecords2() (result []model.Itemapiobindata, err error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/ItemBinCodeAPI"
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

	var autoclik model.Itemapiobindata
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return nil, err
	}

	return []model.Itemapiobindata{autoclik}, nil
}

func CreateNewAutoclik_data2(autoclik_data model.Item_Autoclik_Bin_Code) (value *model.Item_Autoclik_Bin_Code, Error error) {
	db := database.DB
	autoclik_data.Autoclik_binID = uuid.New().String()
	tx := db.Create(&autoclik_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_data, nil
}

func GetAllAutoclik_data() (result []*model.Item_Autoclik_Bin_Code, err error) {
	cachedData, err := redis.GetItem_Autoclik_Bin_Codecache("all_bin_autoclik")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		return cachedData, nil
	}

	db := database.DB
	var records []*model.Item_Autoclik_Bin_Code
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(records) == 0 {
		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
		return records, nil
	}

	err = redis.SetItem_Autoclik_Bin_Codecache("all_bin_autoclik", records)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}

	return records, nil
}

func CountLatestAutoclik_data() (count int, err error) {

	cachedData, err := redis.GetItem_Autoclik_Bin_Codecache("all_bin_autoclik")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		return len(cachedData), nil
	}

	latestAutoclik_data, err := GetAllAutoclik_data()
	if err != nil {
		return 0, err
	}
	err = redis.SetItem_Autoclik_Bin_Codecache("all_bin_autoclik", latestAutoclik_data)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}
	return len(latestAutoclik_data), nil
}

func GetLocation_Code() (result []string, err error) {
	db := database.DB
	var records []*model.Item_Autoclik_Bin_Code
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	seen := make(map[string]bool)

	for _, record := range records {
		if record.Location_Code != nil {
			if _, ok := seen[*record.Location_Code]; !ok {
				result = append(result, *record.Location_Code)
				seen[*record.Location_Code] = true
			}
		}
	}

	// เรียงลำดับ Location_Code ตามลำดับ 1-15
	sort.Strings(result)

	return result, nil
}

func GetByItem_Autoclik_BinNoDB(No string) ([]*model.Item_Autoclik_Bin_Code, error) {
	db := database.DB
	var autoclik []*model.Item_Autoclik_Bin_Code
	tx := db.Preload(clause.Associations).Where("item_no = ?", No).Find(&autoclik)
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

func GetAllRecordsAndCreate2() error {
	values, err := GetAllRecords2()
	if err != nil {
		return err
	}

	var autoclikData []model.Item_Autoclik_Bin_Code
	for _, value := range values {
		for _, item := range value.Value {
			autoclikData = append(autoclikData, model.Item_Autoclik_Bin_Code{
				Autoclik_binID:       uuid.New().String(), // Generate UUID for each record
				Odata_etag:           item.Odata_etag,
				Location_Code:        item.Location_Code,
				Bin_Code:             item.Bin_Code,
				Item_No:              item.Item_No,
				Variant_Code:         item.Variant_Code,
				Unit_of_Measure_Code: item.Unit_of_Measure_Code,
				ItemDesc:             item.ItemDesc,
				ItemDesc2:            item.ItemDesc2,
				ItemCategory:         item.ItemCategory,
				ProductGroup:         item.ProductGroup,
				Fixed:                item.Fixed,
				Default:              item.Default,
				Dedicated:            item.Dedicated,
				CalcQtyUOM:           item.CalcQtyUOM,
				Quantity_Base:        item.Quantity_Base,
				Bin_Type_Code:        item.Bin_Type_Code,
				Zone_Code:            item.Zone_Code,
				Lot_No_Filter:        item.Lot_No_Filter,
				Serial_No_Filter:     item.Serial_No_Filter,
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

func ClearAllItemBin() error {
	db := database.DB
	if err := db.Where("true").Delete(&model.Item_Autoclik_Bin_Code{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAndFetchAllItemBin() error {
	// เคลียร์ข้อมูลทั้งหมดจากตาราง
	if err := ClearAllItemBin(); err != nil {
		return err
	}

	// สร้างข้อมูลใหม่จากข้อมูลที่มีอยู่
	if err := GetAllRecordsAndCreate2(); err != nil {
		return err
	}

	_, err := UpdateRequest_Update_Autoclik_Item_BinData()
	if err != nil {
		return err
	}

	return nil
}

func UpdateRequest_Update_Autoclik_Item_BinData() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "03ec6c6e-8c24-4d3e-a955-331d8e62e21f"
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

func GetChoosebytype() (ItemCategory []string, ProductGroup []string, Location_Code []string, err error) {
	db := database.DB
	var autoclik []*model.Item_Autoclik_Bin_Code
	tx := db.Select("ItemCategory", "ProductGroup", "Location_Code").Preload(clause.Associations).Find(&autoclik)
	if tx.Error != nil {
		return nil, nil, nil, tx.Error
	}

	ItemCategoryMap := make(map[string]struct{})
	ProductGroupMap := make(map[string]struct{})
	Location_CodeMap := make(map[string]struct{})
	for _, data := range autoclik {
		addToMapIfNotNil(ItemCategoryMap, data.ItemCategory)
		addToMapIfNotNil(ProductGroupMap, data.ProductGroup)
		addToMapIfNotNil(Location_CodeMap, data.Location_Code)
	}

	ItemCategory = mapToSortedSlice(ItemCategoryMap)
	ProductGroup = mapToSortedSlice(ProductGroupMap)
	Location_Code = mapToSortedSlice(Location_CodeMap)

	return ItemCategory, ProductGroup, Location_Code, nil
}
