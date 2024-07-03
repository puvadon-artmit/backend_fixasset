package Maliwan_dataServices

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

var Client *http.Client

func GetAllRecordsAndCreate() ([]model.Maliwan_data, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := []string{"(Head Office)"}
	var allResults []model.Maliwan_data

	for _, branchMaliwanCount := range branchMaliwanCounts {
		var result []model.Maliwan_data // Initialize result as nil
		fmt.Println(branchMaliwanCount, "ค่าที่รับมา")
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(branchMaliwanCount))
		ItemApi := "/ItemsAPI"

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
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
				resp.Body.Close()
				break
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return nil, err
			}

			var records model.ItemaMaliwan_data
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

			maliwan_data := model.Maliwan_data{
				Odata_etag:                 value.Odata_etag,
				No:                         value.No,
				No_2:                       value.No_2,
				Description:                value.Description,
				Description_2:              value.Description_2,
				Description_ENG:            value.Description_ENG,
				Inventory:                  value.Inventory,
				Maximum_Inventory:          value.Safety_Stock_Quantity,
				Product_Group_Code:         value.Product_Group_Code,
				Created_From_Nonstock_Item: value.Created_From_Nonstock_Item,
				Line_Type:                  value.Line_Type,
				Type_Insurance:             value.Type_Insurance,
				Type_Registration_fee:      value.Type_Registration_fee,
				Type_Accessory:             value.Type_Accessory,
				Type_Deposit:               value.Type_Deposit,
				Stockkeeping_Unit_Exists:   value.Stockkeeping_Unit_Exists,
				Assembly_BOM:               value.Assembly_BOM,
				Base_Unit_of_Measure:       value.Base_Unit_of_Measure,
				Shelf_No:                   value.Shelf_No,
				Gen_Prod_Posting_Group:     value.Gen_Prod_Posting_Group,
				Inventory_Posting_Group:    value.Inventory_Posting_Group,
				Item_Category_Code:         value.Item_Category_Code,
				VAT_Prod_Posting_Group:     value.VAT_Prod_Posting_Group,
				Item_Disc_Group:            value.Item_Disc_Group,
				Vendor_No:                  value.Vendor_No,
				Vendor_Item_No:             value.Vendor_Item_No,
				Tariff_No:                  value.Tariff_No,
				Search_Description:         value.Search_Description,
				Overhead_Rate:              value.Overhead_Rate,
				Blocked:                    value.Blocked,
				Last_Date_Modified:         value.Last_Date_Modified,
				Sales_Unit_of_Measure:      value.Sales_Unit_of_Measure,
				Replenishment_System:       value.Replenishment_System,
				Purch_Unit_of_Measure:      value.Purch_Unit_of_Measure,
				Lead_Time_Calculation:      value.Lead_Time_Calculation,
				Manufacturing_Policy:       value.Manufacturing_Policy,
				Flushing_Method:            value.Flushing_Method,
				Assembly_Policy:            value.Assembly_Policy,
				Item_Tracking_Code:         value.Item_Tracking_Code,
				Promotion_Model_Code:       value.Promotion_Model_Code,
				Default_Location_Code:      value.Default_Location_Code,
				Qty_on_Purch_Order:         value.Qty_on_Purch_Order,
				Qty_on_Service_Order:       value.Qty_on_Service_Order,
				Qty_on_Sales_Order:         value.Qty_on_Sales_Order,
				Unit_Price_Price:           value.Unit_Price_Price,
				Brand:                      value.Brand,
				Costing_Method:             value.Costing_Method,
				Stockout_Warning:           value.Stockout_Warning,
				Global_Dimension_1_Filter:  value.Global_Dimension_1_Filter,
				Global_Dimension_2_Filter:  value.Global_Dimension_2_Filter,
				Location_Filter:            value.Location_Filter,
				Drop_Shipment_Filter:       value.Drop_Shipment_Filter,
				Variant_Filter:             value.Variant_Filter,
				Lot_No_Filter:              value.Lot_No_Filter,
				Serial_No_Filter:           value.Serial_No_Filter,
				Date_Filter:                value.Date_Filter,
				Branch:                     &branchName,
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

		// Create a new UUID for each record in the chunk
		for j := range chunk {
			chunk[j].Maliwan_dataID = uuid.New().String()
		}

		// Insert the chunk into the database
		db := database.DB
		tx := db.Create(&chunk)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	fmt.Println(len(allResults))
	return allResults, nil
}

func ClearHeadOffice() error {
	db := database.DB
	if err := db.Where("branch = ?", "Head Office").Delete(&model.Maliwan_data{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequest_Update_Maliwan_Branch_HeadOffice() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "bbe0655e-6f30-4861-acdb-9fba13ae08de"
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

func ClearAndFetchHeadOfficeItemBin() error {

	if err := ClearHeadOffice(); err != nil {
		return err
	}

	if err := redis.DeleteKey("Head Office"); err != nil {
		return err
	}

	if err := redis.DeleteKey("branchs"); err != nil {
		return err
	}

	_, err := GetAllRecordsAndCreate()
	if err != nil {
		return err
	}

	_, err = UpdateRequest_Update_Maliwan_Branch_HeadOffice()
	if err != nil {
		return err
	}

	return nil
}

// ------------------------------ ดูดข้อมูล maliwan ทั้งหมด ---------------------------------------------------------------------------

func ClearAllItemBin() error {
	db := database.DB
	if err := db.Where("true").Delete(&model.Maliwan_data{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAndFetchAllItemBin() error {
	// เคลียร์ข้อมูลทั้งหมดจากตาราง
	if err := ClearAllItemBin(); err != nil {
		return err
	}

	// if err := GetAllRecordsAndCreate(); err != nil {
	// 	return err
	// }

	// if err := GetAllRecordsBurirumAndCreate(); err != nil {
	// 	return err
	// }

	// if err := GetAllRecordsNakaAndCreate(); err != nil {
	// 	return err
	// }

	// if err := GetAllRecordsMueang_KrabiAndCreate(); err != nil {
	// 	return err
	// }

	// if err := GetAllRecordsSurinAndCreate(); err != nil {
	// 	return err
	// }

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------------------

// func GetAllMaliwan() (result []*model.Maliwan_data, err error) {
// 	cachedData, err := redis.Getcachemaliwan("maliwan")
// 	if err == nil {
// 		fmt.Println("Successfully retrieved data from cache.")
// 		return cachedData, nil
// 	}

// 	db := database.DB
// 	var autoclik []*model.Maliwan_data
// 	tx := db.Preload(clause.Associations).Find(&autoclik)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	if len(autoclik) == 0 {
// 		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
// 		return autoclik, nil
// 	}

// 	err = redis.Setcachemaliwan("maliwan", autoclik)
// 	if err != nil {
// 		fmt.Println("Warning: Failed to set data in cache:", err)
// 	}

// 	return autoclik, nil
// }

func GetByNoDB(No string) ([]*model.Maliwan_data, error) {
	db := database.DB
	var autoclik []*model.Maliwan_data
	tx := db.Preload(clause.Associations).Where("no = ?", No).Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autoclik, nil
}

func GetAllMaliwan(branchs string) (result []*model.Maliwan_data, err error) {

	cachedData, err := redis.Getcachemaliwan(branchs)
	if err == nil {
		fmt.Println("Successfully retrieved data from cache.")
		return cachedData, nil
	}

	db := database.DB
	var autoclik []*model.Maliwan_data
	tx := db.Select("*").Where("Branch = ? AND Gen_Prod_Posting_Group = ?", branchs, "PARTS").Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}

	err = redis.Setcachemaliwan(branchs, autoclik)
	if err != nil {
		fmt.Println("Warning: Failed to set data in cache:", err)
	}

	return autoclik, nil
}

func GetByCountMaliwan(branchs string) (result []*model.Maliwan_data, err error) {

	db := database.DB
	var autoclik []*model.Maliwan_data
	tx := db.Select("*").Where("Branch = ? AND Gen_Prod_Posting_Group = ?", branchs, "PARTS").Find(&autoclik)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return autoclik, nil
}

func GetByCountBranch(branchs string) (int, error) {
	count, err := GetByCountMaliwan(branchs)
	if err != nil {
		return 0, err
	}
	return len(count), nil
}

// func filterMaliwanByGenProdPostingGroup(data []*model.Maliwan_data) []*model.Maliwan_data {
// 	var filteredData []*model.Maliwan_data
// 	for _, item := range data {
// 		if *item.Gen_Prod_Posting_Group == "PARTS" {
// 			filteredData = append(filteredData, item)
// 		}
// 	}
// 	return filteredData
// }

// func CountLatestMaliwan() (count int, err error) {
// 	maliwan, err := GetAllMaliwan()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return len(maliwan), nil
// }

// func GetByBranch(Maliwan_data string, Limit, Offset int) ([]*model.Maliwan_data, error) {
// 	db := database.DB
// 	var maliwan []*model.Maliwan_data
// 	tx := db.Raw("SELECT * FROM maliwan_data WHERE Branch = ? AND Gen_Prod_Posting_Group = 'PARTS' LIMIT ? OFFSET ? ", Maliwan_data, Limit, Offset).Scan(&maliwan)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return maliwan, nil
// }

func GetByBranch(branch string, limit, offset int) ([]*model.Maliwan_data, error) {

	var maliwan []*model.Maliwan_data
	db := database.DB

	if err := db.Select("no, description, unit_price_price, last_date_modified, branch").
		Where("Branch = ? AND Gen_Prod_Posting_Group = ?", branch, "PARTS").
		Limit(limit).
		Offset(offset).
		Find(&maliwan).Error; err != nil {
		return nil, err
	}

	return maliwan, nil
}

func GetAllMaliwanData() (result []*model.Maliwan_data, err error) {

	// cachedData, err := redis.Getcachemaliwan("maliwan-alldata")
	// if err == nil {
	// 	fmt.Println("Successfully retrieved data from cache.")
	// 	return cachedData, nil
	// }

	db := database.DB
	var records []*model.Maliwan_data
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// err = redis.Setcachemaliwan("maliwan-alldata", records)
	// if err != nil {
	// 	fmt.Println("Warning: Failed to set data in cache:", err)
	// }

	return records, nil
}

func CountAllMaliwan() (count int, err error) {
	db := database.DB
	var totalCount int64
	tx := db.Model(&model.Maliwan_data{}).Count(&totalCount)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(totalCount), nil
}

func CountMaliwanDataByBranch() (map[string]int, error) {

	db := database.DB
	branchCounts := make(map[string]int)

	var records []*model.Maliwan_data
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}

	for _, record := range records {
		branchCounts[*record.Branch]++
	}

	return branchCounts, nil
}

func GetByItem_Category_CodeDB(Item_category_code string) ([]*model.Maliwan_data, error) {

	db := database.DB
	var item_category_code []*model.Maliwan_data
	tx := db.Preload(clause.Associations).Where("Item_Category_Code = ?", Item_category_code).Find(&item_category_code)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return item_category_code, nil
}

// func GetByBranch(branch string, limit, offset int) ([]*model.Maliwan_data, error) {
// 	cachedMaliwan, err := redis.Getcachemaliwan(branch)
// 	if err == nil {
// 		fmt.Println("ข้อมูลถูกดึงจากแคช")
// 		return cachedMaliwan, nil
// 	}

// 	var maliwan []*model.Maliwan_data
// 	db := database.DB

// 	if err := db.Select("no, description, unit_price_price, last_date_modified, branch").
// 		Where("Branch = ? AND Gen_Prod_Posting_Group = ?", branch, "PARTS").
// 		Limit(limit).
// 		Offset(offset).
// 		Find(&maliwan).Error; err != nil {
// 		return nil, err
// 	}

// 	err = redis.Lpushmaliwan(branch, maliwan)
// 	if err != nil {
// 		fmt.Println("เกิดข้อผิดพลาดในการบันทึกข้อมูลลงใน Redis List:", err)
// 	}

// 	return maliwan, nil
// }

// func fetchDataFromDatabase(Branch string) (int, error) {
// 	db := database.DB
// 	var count int64
// 	tx := db.Model(&model.Maliwan_data{}).
// 		Where("Branch = ?", Branch).
// 		Where("Gen_Prod_Posting_Group = ?", "PARTS").
// 		Count(&count)
// 	if tx.Error != nil {
// 		return 0, tx.Error
// 	}

// 	err := redis.Setcachemaliwan(Branch, nil)
// 	if err != nil {
// 		fmt.Println("Warning: Failed to set data in cache:", err)
// 	}

// 	return int(count), nil
// }

// func GetByCountBranch(Branch string) (int, error) {
// 	cachedData, err := redis.Getcachemaliwan(Branch)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return len(cachedData), nil
// }

// func GetByCountBranch(Maliwan_data string) ([]*model.Maliwan_data, error) {
// 	cachedData, err := redis.Getcachemaliwan(Maliwan_data)
// 	if err == nil {
// 		fmt.Println("Successfully retrieved data from cache by branch.")
// 		fmt.Println(len(cachedData))
// 		return filterMaliwanByGenProdPostingGroup(cachedData), nil
// 	}

// 	db := database.DB
// 	var maliwan []*model.Maliwan_data
// 	tx := db.Preload(clause.Associations).Where("Branch = ?", Maliwan_data).Find(&maliwan)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	if len(maliwan) == 0 {
// 		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
// 		return maliwan, nil
// 	}

// 	err = redis.Setcachemaliwan(Maliwan_data, maliwan)
// 	if err != nil {
// 		fmt.Println("Warning: Failed to set data in cache:", err)
// 	}

// 	return cachedData, nil
// }

// func GetByBranch(Maliwan_data string) ([]*model.Maliwan_data, error) {

// 	db := database.DB
// 	var maliwan []*model.Maliwan_data
// 	tx := db.Preload(clause.Associations).Where("Branch = ?", Maliwan_data).Find(&maliwan)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	if len(maliwan) == 0 {
// 		fmt.Println("ไม่พบข้อมูลในแคช กำลังดึงข้อมูลจากฐานข้อมูล...")
// 		return maliwan, nil
// 	}

// 	return maliwan, nil
// }

// -------------------------------------------------------------------------------------------------------------------------------------------

func GetFilterRecords() (result []model.Itemapivalue, err error) {

	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}

	var records model.Itemapiodata

	ArrayofCompanyName := [...]string{"(Mueang Krabi)", "(Naka)", "(Head Office)", "(Burirum)", "(Surin)"}
	for _, company := range ArrayofCompanyName {
		fmt.Println(company)
		top := 7000
		skip := 0
		NAV_SERVER_HMW := "http://10.0.1.2:4048/API-ACGHMW/ODataV4"
		Company := fmt.Sprintf("/Company('%v%v')", url.PathEscape("Honda Maliwan "), url.PathEscape(company))
		ItemApi := "/ItemsAPI"

		param := fmt.Sprintf("?$top=%v&$skip=%v", top, skip)
		path_url2 := NAV_SERVER_HMW + Company + ItemApi + param
		req, err := http.NewRequest("GET", path_url2, strings.NewReader(""))

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
			err = json.Unmarshal(body, &records)
			if err != nil {
				return nil, err
			}

			result = append(result, records.Value...)

			if len(records.Value) < top {
				break
			}
			skip += top

			param = fmt.Sprintf("?$top=%v&$skip=%v&", top, skip)
			path_url2 = NAV_SERVER_HMW + Company + ItemApi + param
			req, err = http.NewRequest("GET", path_url2, strings.NewReader(""))
			if err != nil {
				return nil, err
			}
		}
	}

	fmt.Println(len(result))
	return result, nil
}

func CreatItem_Category_Code(maliwan_data model.Maliwan_Item_Category_Code) (value *model.Maliwan_Item_Category_Code, Error error) {
	db := database.DB
	maliwan_data.Item_Category_CodeID = uuid.New().String()
	tx := db.Create(&maliwan_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_data, nil
}

func GetAllRecordsAndCreateItem_Category_Code() error {
	values, err := GetFilterRecords()
	if err != nil {
		return err
	}

	for _, value := range values {
		// Check if the Item_Category_Code already exists in the database
		exists, err := ItemCategoryCodeExists(value.Item_Category_Code)
		if err != nil {
			return err
		}

		// If it doesn't exist, create a new record
		if !exists {
			maliwan_data := model.Maliwan_Item_Category_Code{
				Item_Category_Code: &value.Item_Category_Code,
			}

			_, err := CreatItem_Category_Code(maliwan_data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ItemCategoryCodeExists checks if an Item_Category_Code exists in the database
func ItemCategoryCodeExists(itemCategoryCode string) (bool, error) {
	db := database.DB
	var count int64
	tx := db.Model(&model.Maliwan_Item_Category_Code{}).Where("item_category_code = ?", itemCategoryCode).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}

func GetAllMaliwan_Item_Category_Code() (result []*model.Maliwan_Item_Category_Code, err error) {
	db := database.DB
	var records []*model.Maliwan_Item_Category_Code
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func GetAllMaliwan_Item_Category_Code_And_Filter(ctx context.Context) (itemCategoryCodes, branches []string, err error) {
	db := database.DB

	// Step 1: Get Branches from cache or database
	cachedBranch, err := redis.GetCacheFilter(ctx, "Branch")
	if err != nil {
		return nil, nil, err
	}
	if cachedBranch != nil {
		branches = cachedBranch
	} else {
		var records []*model.Maliwan_data
		tx := db.Preload(clause.Associations).Find(&records)
		if tx.Error != nil {
			return nil, nil, tx.Error
		}

		seenBranch := make(map[string]bool)
		for _, record := range records {
			if record.Branch != nil && *record.Branch != "" {
				if _, ok := seenBranch[*record.Branch]; !ok {
					branches = append(branches, *record.Branch)
					seenBranch[*record.Branch] = true
				}
			}
		}

		sort.Strings(branches)

		err := redis.SetCacheFilter(ctx, "Branch", branches)
		if err != nil {
			fmt.Println("Error setting Branch cache:", err)
		}
	}

	// Step 2: Get Item_Category_Code from database
	var itemCategoryRecords []*model.Maliwan_Item_Category_Code
	tx := db.Preload(clause.Associations).Find(&itemCategoryRecords)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}

	for _, record := range itemCategoryRecords {
		if record.Item_Category_Code != nil {
			itemCategoryCodes = append(itemCategoryCodes, *record.Item_Category_Code)
		}
	}

	return itemCategoryCodes, branches, nil
}

func CreateNewMaliwan_Update_Story(inspection_story model.Maliwan_Update_Story) (value *model.Maliwan_Update_Story, Error error) {
	db := database.DB
	inspection_story.Maliwan_Update_StoryID = uuid.New().String()
	group_api := "ทุกสาขา"
	inspection_story.Group_api = group_api

	maliwan_update_name := "อัพเดทข้อมูล"
	inspection_story.Maliwan_Update_Name = maliwan_update_name

	user_id := "86d233a4-b45e-4a2f-940e-a17eafb79234"
	inspection_story.UserID = user_id

	tx := db.Create(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func GetMaliwanDataDB(limit int, offset int) ([]*model.Maliwan_data, error) {
	db := database.DB

	var records []*model.Maliwan_data
	tx := db.Preload(clause.Associations).Limit(limit).Offset(offset).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var jsonData []*model.Maliwan_data
	for _, record := range records {
		jsonData = append(jsonData, &model.Maliwan_data{
			No:                record.No,
			Description:       record.Description,
			Branch:            record.Branch,
			Inventory:         record.Inventory,
			Maximum_Inventory: record.Maximum_Inventory,
		})
	}
	return jsonData, nil
}

func GetTotalMaliwanData() (int64, error) {
	var count int64
	db := database.DB
	tx := db.Model(&model.Maliwan_data{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return count, nil
}

func GetAllbranchsMaliwan(branchs string, gen_prod_posting_group string) ([]*model.Maliwan_data, error) {
	db := database.DB
	var maliwan []*model.Maliwan_data
	tx := db.Preload(clause.Associations).Where("branch = ? AND gen_prod_posting_group = ?", branchs, gen_prod_posting_group).Find(&maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan, nil
}

func GetBatchOfBranchsMaliwan(Branch string, GenProdPostingGroup string, offset int, limit int) ([]*model.Maliwan_data, error) {
	db := database.DB
	var maliwan []*model.Maliwan_data
	tx := db.Preload(clause.Associations).Where("branch = ? AND gen_prod_posting_group = ?", Branch, GenProdPostingGroup).Offset(offset).Limit(limit).Find(&maliwan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return maliwan, nil
}
