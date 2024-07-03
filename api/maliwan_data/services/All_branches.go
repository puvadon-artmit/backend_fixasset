package Maliwan_dataServices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/spf13/viper"
	httpntlm "github.com/vadimi/go-http-ntlm/v2"
	"gorm.io/gorm/clause"
)

func GetAllRecordsBurirumAndCreate() ([]model.Maliwan_data, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := []string{"(Burirum)"}
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

func UpdateRequest_Update_Maliwan_Branch_Burirum() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "01d59511-d1fd-41f6-bfba-39261a88b386"
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

func ClearBurirumItemBin() error {
	db := database.DB
	if err := db.Where("branch = ?", "Burirum").Delete(&model.Maliwan_data{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAndFetchBurirumItemBin() error {

	if err := ClearBurirumItemBin(); err != nil {
		return err
	}

	_, err := GetAllRecordsBurirumAndCreate()
	if err != nil {
		return err
	}

	_, err = UpdateRequest_Update_Maliwan_Branch_Burirum()
	if err != nil {
		return err
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------

func GetAllRecordsNakaAndCreate() ([]model.Maliwan_data, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := []string{"(Naka)"}
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

func ClearNakaItemBin() error {
	db := database.DB
	if err := db.Where("branch = ?", "Naka").Delete(&model.Maliwan_data{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequest_Update_Maliwan_Branch_Naka() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "6b1c8f80-5a33-48aa-8ed3-fbd30a17f07a"
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

func ClearAndFetchNakaItemBin() error {

	if err := ClearNakaItemBin(); err != nil {
		return err
	}

	_, err := GetAllRecordsNakaAndCreate()
	if err != nil {
		return err
	}

	_, err = UpdateRequest_Update_Maliwan_Branch_Naka()
	if err != nil {
		return err
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------

func GetAllRecordsMueang_KrabiAndCreate() ([]model.Maliwan_data, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := []string{"(Mueang Krabi)"}
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

func ClearMueang_KrabiItemBin() error {
	db := database.DB
	if err := db.Where("branch = ?", "Mueang Krabi").Delete(&model.Maliwan_data{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequest_Update_Maliwan_Branch_Mueang_Krabi() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "fe90528b-ed98-4fbe-a50c-ecb8d9a7f144"
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

func ClearAndFetchMueang_KrabiItemBin() error {

	if err := ClearMueang_KrabiItemBin(); err != nil {
		return err
	}

	_, err := GetAllRecordsMueang_KrabiAndCreate()
	if err != nil {
		return err
	}

	_, err = UpdateRequest_Update_Maliwan_Branch_Mueang_Krabi()
	if err != nil {
		return err
	}

	return nil
}

// ------------------------------------------------------------------------------------------------------------------

func GetAllRecordsSurinAndCreate() ([]model.Maliwan_data, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := []string{"(Surin)"}
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

func ClearSurinItemBin() error {
	db := database.DB
	if err := db.Where("branch = ?", "Surin").Delete(&model.Maliwan_data{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRequest_Update_Maliwan_Branch_Surin() (*model.Request_Update_Data, error) {
	db := database.DB

	id := "c738c18e-daa2-420d-9bea-a6cfe83d3c69"
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

func ClearAndFetchSurinItemBin() error {

	if err := ClearSurinItemBin(); err != nil {
		return err
	}

	_, err := GetAllRecordsSurinAndCreate()
	if err != nil {
		return err
	}

	_, err = UpdateRequest_Update_Maliwan_Branch_Surin()
	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------------------------------------

func ClearAndPullMaliwan_Item() error {
	if err := DeleteAllDataMaliwan(); err != nil {
		return err
	}

	time.Sleep(1 * time.Minute)
	// ดึงข้อมูลทั้งหมดโดยไม่ต้องส่ง `fiber.Ctx`
	data, err := GetAllDataRecords()
	if err != nil {
		return err
	}

	fmt.Printf("ดึงข้อมูลได้ทั้งหมด %d รายการ\n", len(data))
	return nil
}

func DeleteAllDataMaliwan() error {
	db := database.DB
	result := db.Exec("DELETE FROM maliwan_data")
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("ลบข้อมูลทั้งหมดเสร็จสิ้น")
	return nil
}

func GetAllDataRecords() ([]model.Maliwan_data, error) {
	Client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   viper.GetString("nav.domain"),
			User:     viper.GetString("nav.username"),
			Password: viper.GetString("nav.password"),
		},
	}
	branchMaliwanCounts := []string{"(Head Office)", "(Mueang Krabi)", "(Naka)", "(Burirum)", "(Surin)"}
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

func CreateMaliwandata(maliwan_data model.Maliwan_data) (value *model.Maliwan_data, Error error) {
	db := database.DB
	maliwan_data.Maliwan_dataID = uuid.New().String()
	tx := db.Create(&maliwan_data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &maliwan_data, nil
}
