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

// var Client *http.Client

// func GetAllRecords() (result []model.Itemapivalue, err error) {

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

func GetAllRecords() (result []model.Item_AutoclikPRD, err error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	NAV_SERVER_AUTOCLIK := "http://10.0.1.2:3048/API-CLIK/ODataV4/Company('Autoclik')"
	ItemAPI := "/ItemAPI"
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

	var autoclik model.Item_AutoclikPRD
	err = json.Unmarshal(body, &autoclik)
	if err != nil {
		return nil, err
	}

	return []model.Item_AutoclikPRD{autoclik}, nil
}

func CreateNewAutoclik_data(autoclik_data model.Item_Autoclik) (value *model.Item_Autoclik, Error error) {
	db := database.DB
	autoclik_data.Autoclik_dataID = uuid.New().String()
	tx := db.Create(&autoclik_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik_data, nil
}

func GetAllRecordsAndCreate() error {
	values, err := GetAllRecords()
	if err != nil {
		return err
	}

	var autoclikData []model.Item_Autoclik

	for _, value := range values {
		for _, item := range value.Value {
			autoclikData = append(autoclikData, model.Item_Autoclik{
				Autoclik_dataID:            uuid.New().String(),
				No:                         item.No,
				No_2:                       item.No_2,
				Description:                item.Description,
				Description_2:              item.Description_2,
				Description_ENG:            item.Description_ENG,
				Inventory:                  item.Inventory,
				Maximum_Inventory:          item.Safety_Stock_Quantity,
				Product_Group_Code:         item.Product_Group_Code,
				Created_From_Nonstock_Item: item.Created_From_Nonstock_Item,
				Line_Type:                  item.Line_Type,
				Type_Insurance:             item.Type_Insurance,
				Type_Registration_fee:      item.Type_Registration_fee,
				Type_Accessory:             item.Type_Accessory,
				Type_Deposit:               item.Type_Deposit,
				Stockkeeping_Unit_Exists:   item.Stockkeeping_Unit_Exists,
				Assembly_BOM:               item.Assembly_BOM,
				Base_Unit_of_Measure:       item.Base_Unit_of_Measure,
				Shelf_No:                   item.Shelf_No,
				Gen_Prod_Posting_Group:     item.Gen_Prod_Posting_Group,
				Inventory_Posting_Group:    item.Inventory_Posting_Group,
				Item_Category_Code:         item.Item_Category_Code,
				VAT_Prod_Posting_Group:     item.VAT_Prod_Posting_Group,
				Item_Disc_Group:            item.Item_Disc_Group,
				Vendor_No:                  item.Vendor_No,
				Vendor_Item_No:             item.Vendor_Item_No,
				Tariff_No:                  item.Tariff_No,
				Search_Description:         item.Search_Description,
				Overhead_Rate:              item.Overhead_Rate,
				Blocked:                    item.Blocked,
				Last_Date_Modified:         item.Last_Date_Modified,
				Sales_Unit_of_Measure:      item.Sales_Unit_of_Measure,
				Replenishment_System:       item.Replenishment_System,
				Purch_Unit_of_Measure:      item.Purch_Unit_of_Measure,
				Lead_Time_Calculation:      item.Lead_Time_Calculation,
				Manufacturing_Policy:       item.Manufacturing_Policy,
				Flushing_Method:            item.Flushing_Method,
				Assembly_Policy:            item.Assembly_Policy,
				Item_Tracking_Code:         item.Item_Tracking_Code,
				Promotion_Model_Code:       item.Promotion_Model_Code,
				Default_Location_Code:      item.Default_Location_Code,
				Qty_on_Purch_Order:         item.Qty_on_Purch_Order,
				Qty_on_Service_Order:       item.Qty_on_Service_Order,
				Qty_on_Sales_Order:         item.Qty_on_Sales_Order,
				Unit_Price_Price:           item.Unit_Price_Price,
				Brand:                      item.Brand,
				Costing_Method:             item.Costing_Method,
				Stockout_Warning:           item.Stockout_Warning,
				Global_Dimension_1_Filter:  item.Global_Dimension_1_Filter,
				Global_Dimension_2_Filter:  item.Global_Dimension_2_Filter,
				Location_Filter:            item.Location_Filter,
				Drop_Shipment_Filter:       item.Drop_Shipment_Filter,
				Variant_Filter:             item.Variant_Filter,
				Lot_No_Filter:              item.Lot_No_Filter,
				Serial_No_Filter:           item.Serial_No_Filter,
				Date_Filter:                item.Date_Filter,
			})
		}
	}

	// แบ่งข้อมูลเป็นชุดเพื่อ insert ลงในฐานข้อมูล
	chunkSize := 1000
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

func GetById(id string) (value *model.Item_Autoclik, Error error) {
	db := database.DB
	autoclik := model.Item_Autoclik{Autoclik_dataID: id}
	tx := db.Preload(clause.Associations).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik, nil
}

func GetByNoDB(No string) (*model.Item_Autoclik, error) {
	db := database.DB
	autoclik := model.Item_Autoclik{}
	tx := db.Preload(clause.Associations).Where("no = ?", No).First(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &autoclik, nil
}

func GetBygen_prod_posting_groupDB(No string) ([]*model.Item_Autoclik, error) {
	db := database.DB
	autocliks := []*model.Item_Autoclik{}
	tx := db.Preload(clause.Associations).Where("gen_prod_posting_group = ?", No).Find(&autocliks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autocliks, nil
}

func GetAllAutoclik() (result []*model.Item_Autoclik, err error) {
	cachedData, err := redis.Getcache("all_autoclik")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		return cachedData, nil
	}

	db := database.DB
	var autoclik []*model.Item_Autoclik
	tx := db.Preload(clause.Associations).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(autoclik) == 0 {
		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
		return autoclik, nil
	}

	err = redis.Setcache("all_autoclik", autoclik)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}

	return autoclik, nil
}

func GetAllAutoclikLimit(limit, offset int) (result []*model.Item_Autoclik, err error) {
	cachedData, err := redis.Getcache("all_autocliks")
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		// Return only the subset of cachedData based on limit and offset
		if offset >= len(cachedData) {
			return []*model.Item_Autoclik{}, nil
		}
		endIndex := offset + limit
		if endIndex > len(cachedData) {
			endIndex = len(cachedData)
		}
		return cachedData[offset:endIndex], nil
	}

	db := database.DB
	var autoclik []*model.Item_Autoclik
	tx := db.Preload(clause.Associations).Limit(limit).Offset(offset).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(autoclik) == 0 {
		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
		return autoclik, nil
	}

	err = redis.Setcache("all_autocliks", autoclik)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}

	return autoclik, nil
}

// func GetByCount() (int, error) {
// 	allAutoclik, err := GetAllAutoclik()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return len(allAutoclik), nil
// }

func GetByCount() (int, error) {
	cachedCount, err := redis.Getcache("autoclik_counts")
	if err == nil {
		fmt.Println("Successfully retrieved count from cache.")
		return len(cachedCount), nil
	}

	allAutoclik, err := GetAllAutoclik()
	if err != nil {
		return 0, err
	}
	count := len(allAutoclik)

	var autoclikSlice []*model.Item_Autoclik
	for i := 0; i < count; i++ {
		autoclikSlice = append(autoclikSlice, &model.Item_Autoclik{})
	}

	err = redis.Setcache("autoclik_counts", autoclikSlice)
	if err != nil {
		fmt.Println("Failed to cache count:", err)
	}

	return count, nil
}

func GetAllCategory() (categories []string, inventoryGroups []string, brands []string, err error) {
	db := database.DB
	var autoclik []*model.Item_Autoclik
	tx := db.Select("Item_Category_Code", "Inventory_Posting_Group", "Brand").Preload(clause.Associations).Find(&autoclik)
	if tx.Error != nil {
		return nil, nil, nil, tx.Error
	}

	categoryMap := make(map[string]struct{})
	inventoryMap := make(map[string]struct{})
	brandMap := make(map[string]struct{})
	for _, data := range autoclik {
		addToMapIfNotNil(categoryMap, data.Item_Category_Code)
		addToMapIfNotNil(inventoryMap, data.Inventory_Posting_Group)
		addToMapIfNotNil(brandMap, data.Brand)
	}

	categories = mapToSortedSlice(categoryMap)
	inventoryGroups = mapToSortedSlice(inventoryMap)
	brands = mapToSortedSlice(brandMap)

	return categories, inventoryGroups, brands, nil
}

func addToMapIfNotNil(m map[string]struct{}, value *string) {
	if value != nil && *value != "" {
		m[*value] = struct{}{}
	}
}

func mapToSortedSlice(m map[string]struct{}) []string {
	slice := make([]string, 0, len(m))
	for key := range m {
		slice = append(slice, key)
	}
	sort.Strings(slice)
	return slice
}

func ClearAllItems() error {
	db := database.DB
	if err := db.Where("true").Delete(&model.Item_Autoclik{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAndFetchAllItems() error {
	// เคลียร์ข้อมูลทั้งหมดจากตาราง
	if err := ClearAllItems(); err != nil {
		return err
	}

	// สร้างข้อมูลใหม่จากข้อมูลที่มีอยู่
	if err := GetAllRecordsAndCreate(); err != nil {
		return err
	}

	_, err := UpdateRequest_Update_AutoclikData()
	if err != nil {
		return err
	}

	return nil
}

func UpdateRequest_Update_AutoclikData() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "9ff54a16-74e8-421e-beb8-3eb1e1ee9b2b"
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
