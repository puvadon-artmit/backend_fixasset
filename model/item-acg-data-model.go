package model

import "time"

type Itemapiodata struct {
	Odata_context string         `json:"@odata.context"`
	Value         []Itemapivalue `json:"value"`
}

type Item_AutoclikPRD struct {
	Odata_context string          `json:"@odata.context"`
	Value         []Item_Autoclik `json:"value"`
}

type Itemapiobindata struct {
	Odata_context string                   `json:"@odata.context"`
	Value         []Item_Autoclik_Bin_Code `json:"value"`
}

type ItemAutoclik_Fixed_Asset struct {
	Odata_context string                 `json:"@odata.context"`
	Value         []Autoclik_Fixed_Asset `json:"value"`
}

type Autoclik_Fixed_Asset struct {
	Autoclik_Fixed_AssetId  string  `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_id"`
	No                      *string `json:"No"`
	Description             *string `json:"Description"`
	Description_2           *string `json:"Description_2"`
	Ref_Code                *string `json:"Ref_Code"`
	Serial_No               *string `json:"Serial_No"`
	Search_Description      *string `json:"Search_Description"`
	Responsible_Employee    *string `json:"Responsible_Employee"`
	FA_Posting_Group        *string `json:"FA_Posting_Group"`
	FA_Class_Code           *string `json:"FA_Class_Code"`
	FA_Subclass_Code        *string `json:"FA_Subclass_Code"`
	FA_Location_Code        *string `json:"FA_Location_Code"`
	FA_Department_Code      *string `json:"FA_Department_Code"`
	Budgeted_Asset          bool    `json:"Budgeted_Asset"`
	Inactive                bool    `json:"Inactive"`
	Blocked                 bool    `json:"Blocked"`
	Global_Dimension_1_Code *string `json:"Global_Dimension_1_Code"`
	Global_Dimension_2_Code *string `json:"Global_Dimension_2_Code"`
	Main_Asset_Component    *string `json:"Main_Asset_Component"`
	Component_of_Main_Asset *string `json:"Component_of_Main_Asset"`
}

type Item_Autoclik_Bin_Code struct {
	Autoclik_binID       string     `gorm:"type:uuid;primaryKey" json:"autoclik_bin_id"`
	Odata_etag           *string    `json:"@odata.etag"`
	Location_Code        *string    `json:"Location_Code"`
	Bin_Code             *string    `json:"Bin_Code"`
	Item_No              *string    `json:"Item_No"`
	Variant_Code         *string    `json:"Variant_Code"`
	Unit_of_Measure_Code *string    `json:"Unit_of_Measure_Code"`
	ItemDesc             *string    `json:"ItemDesc"`
	ItemDesc2            *string    `json:"ItemDesc2"`
	ItemCategory         *string    `json:"ItemCategory"`
	ProductGroup         *string    `json:"ProductGroup"`
	Fixed                bool       `json:"Fixed"`
	Default              bool       `json:"Default"`
	Dedicated            bool       `json:"Dedicated"`
	CalcQtyUOM           float64    `json:"CalcQtyUOM"`
	Quantity_Base        int64      `json:"Quantity_Base"`
	Bin_Type_Code        *string    `json:"Bin_Type_Code"`
	Zone_Code            *string    `json:"Zone_Code"`
	Lot_No_Filter        *string    `json:"Lot_No_Filter"`
	Serial_No_Filter     *string    `json:"Serial_No_Filter"`
	CreatedAt            *time.Time `json:"created_at"`
}

type Itemapivalue struct {
	Odata_etag                 string  `json:"@odata.etag"`
	No                         string  `json:"No"`
	No_2                       string  `json:"No_2"`
	Description                string  `json:"Description"`
	Description_2              string  `json:"Description_2"`
	Description_ENG            string  `json:"Description_ENG"`
	Inventory                  float64 `json:"Inventory"`
	Maximum_Inventory          int32   `json:"Maximum_Inventory"`
	Safety_Stock_Quantity      int32   `json:"Safety_Stock_Quantity"`
	Product_Group_Code         string  `json:"Product_Group_Code"`
	Created_From_Nonstock_Item bool    `json:"Created_From_Nonstock_Item"`
	Line_Type                  string  `json:"Line_Type"`
	Type_Insurance             bool    `json:"Type_Insurance"`
	Type_Registration_fee      bool    `json:"Type_Registration_fee"`
	Type_Accessory             bool    `json:"Type_Accessory"`
	Type_Deposit               bool    `json:"Type_Deposit"`
	Stockkeeping_Unit_Exists   bool    `json:"Stockkeeping_Unit_Exists"`
	Assembly_BOM               bool    `json:"Assembly_BOM"`
	Base_Unit_of_Measure       string  `json:"Base_Unit_of_Measure"`
	Shelf_No                   string  `json:"Shelf_No"`
	Gen_Prod_Posting_Group     string  `json:"Gen_Prod_Posting_Group"`
	Inventory_Posting_Group    string  `json:"Inventory_Posting_Group"`
	Item_Category_Code         string  `json:"Item_Category_Code"`
	VAT_Prod_Posting_Group     string  `json:"VAT_Prod_Posting_Group"`
	Item_Disc_Group            string  `json:"Item_Disc_Group"`
	Vendor_No                  string  `json:"Vendor_No"`
	Vendor_Item_No             string  `json:"Vendor_Item_No"`
	Tariff_No                  string  `json:"Tariff_No"`
	Search_Description         string  `json:"Search_Description"`
	Overhead_Rate              int32   `json:"Overhead_Rate"`
	Blocked                    bool    `json:"Blocked"`
	Last_Date_Modified         string  `json:"Last_Date_Modified"`
	Sales_Unit_of_Measure      string  `json:"Sales_Unit_of_Measure"`
	Replenishment_System       string  `json:"Replenishment_System"`
	Purch_Unit_of_Measure      string  `json:"Purch_Unit_of_Measure"`
	Lead_Time_Calculation      string  `json:"Lead_Time_Calculation"`
	Manufacturing_Policy       string  `json:"Manufacturing_Policy"`
	Flushing_Method            string  `json:"Flushing_Method"`
	Assembly_Policy            string  `json:"Assembly_Policy"`
	Item_Tracking_Code         string  `json:"Item_Tracking_Code"`
	Promotion_Model_Code       string  `json:"Promotion_Model_Code"`
	Default_Location_Code      string  `json:"Default_Location_Code"`
	Qty_on_Purch_Order         int32   `json:"Qty_on_Purch_Order"`
	Qty_on_Service_Order       float64 `json:"Qty_on_Service_Order"`
	Qty_on_Sales_Order         float64 `json:"Qty_on_Sales_Order"`
	Unit_Price_Price           float64 `json:"Unit_Price_Price"`
	Brand                      string  `json:"Brand"`
	Costing_Method             string  `json:"Costing_Method"`
	Stockout_Warning           string  `json:"Stockout_Warning"`
	Global_Dimension_1_Filter  string  `json:"Global_Dimension_1_Filter"`
	Global_Dimension_2_Filter  string  `json:"Global_Dimension_2_Filter"`
	Location_Filter            string  `json:"Location_Filter"`
	Drop_Shipment_Filter       string  `json:"Drop_Shipment_Filter"`
	Variant_Filter             string  `json:"Variant_Filter"`
	Lot_No_Filter              string  `json:"Lot_No_Filter"`
	Serial_No_Filter           string  `json:"Serial_No_Filter"`
	Date_Filter                string  `json:"Date_Filter"`
}

type Item_Autoclik struct {
	Autoclik_dataID string `gorm:"type:uuid;primaryKey" json:"autoclik_data_id"`
	// Odata_etag                 *string  `json:"@odata.etag"`
	No              *string `json:"No"`
	No_2            *string `json:"No_2"`
	Description     *string `json:"Description"`
	Description_2   *string `json:"Description_2"`
	Description_ENG *string `json:"Description_ENG"`
	// Amount                     float32  `sql:"type:decimal(20,2);"`
	Inventory                  *float64   `json:"Inventory"`
	Maximum_Inventory          *int32     `json:"Maximum_Inventory"`
	Safety_Stock_Quantity      *int32     `json:"Safety_Stock_Quantity"`
	Product_Group_Code         *string    `json:"Product_Group_Code"`
	Created_From_Nonstock_Item *bool      `json:"Created_From_Nonstock_Item"`
	Line_Type                  *string    `json:"Line_Type"`
	Type_Insurance             *bool      `json:"Type_Insurance"`
	Type_Registration_fee      *bool      `json:"Type_Registration_fee"`
	Type_Accessory             *bool      `json:"Type_Accessory"`
	Type_Deposit               *bool      `json:"Type_Deposit"`
	Stockkeeping_Unit_Exists   *bool      `json:"Stockkeeping_Unit_Exists"`
	Assembly_BOM               *bool      `json:"Assembly_BOM"`
	Base_Unit_of_Measure       *string    `json:"Base_Unit_of_Measure"`
	Shelf_No                   *string    `json:"Shelf_No"`
	Gen_Prod_Posting_Group     *string    `json:"Gen_Prod_Posting_Group"`
	Inventory_Posting_Group    *string    `json:"Inventory_Posting_Group"`
	Item_Category_Code         *string    `json:"Item_Category_Code"`
	VAT_Prod_Posting_Group     *string    `json:"VAT_Prod_Posting_Group"`
	Item_Disc_Group            *string    `json:"Item_Disc_Group"`
	Vendor_No                  *string    `json:"Vendor_No"`
	Vendor_Item_No             *string    `json:"Vendor_Item_No"`
	Tariff_No                  *string    `json:"Tariff_No"`
	Search_Description         *string    `json:"Search_Description"`
	Overhead_Rate              *int32     `json:"Overhead_Rate"`
	Blocked                    *bool      `json:"Blocked"`
	Last_Date_Modified         *string    `json:"Last_Date_Modified"`
	Sales_Unit_of_Measure      *string    `json:"Sales_Unit_of_Measure"`
	Replenishment_System       *string    `json:"Replenishment_System"`
	Purch_Unit_of_Measure      *string    `json:"Purch_Unit_of_Measure"`
	Lead_Time_Calculation      *string    `json:"Lead_Time_Calculation"`
	Manufacturing_Policy       *string    `json:"Manufacturing_Policy"`
	Flushing_Method            *string    `json:"Flushing_Method"`
	Assembly_Policy            *string    `json:"Assembly_Policy"`
	Item_Tracking_Code         *string    `json:"Item_Tracking_Code"`
	Promotion_Model_Code       *string    `json:"Promotion_Model_Code"`
	Default_Location_Code      *string    `json:"Default_Location_Code"`
	Qty_on_Purch_Order         *int32     `json:"Qty_on_Purch_Order"`
	Qty_on_Service_Order       *float64   `json:"Qty_on_Service_Order"`
	Qty_on_Sales_Order         *float64   `json:"Qty_on_Sales_Order"`
	Unit_Price_Price           *float64   `json:"Unit_Price_Price"`
	Brand                      *string    `json:"Brand"`
	Costing_Method             *string    `json:"Costing_Method"`
	Stockout_Warning           *string    `json:"Stockout_Warning"`
	Global_Dimension_1_Filter  *string    `json:"Global_Dimension_1_Filter"`
	Global_Dimension_2_Filter  *string    `json:"Global_Dimension_2_Filter"`
	Location_Filter            *string    `json:"Location_Filter"`
	Drop_Shipment_Filter       *string    `json:"Drop_Shipment_Filter"`
	Variant_Filter             *string    `json:"Variant_Filter"`
	Lot_No_Filter              *string    `json:"Lot_No_Filter"`
	Serial_No_Filter           *string    `json:"Serial_No_Filter"`
	Date_Filter                *string    `json:"Date_Filter"`
	CreatedAt                  *time.Time `json:"created_at"`
}

type ItemaMaliwan_data struct {
	Odata_context string         `json:"@odata.context"`
	Value         []Maliwan_data `json:"value"`
}

type Maliwan_data struct {
	Maliwan_dataID             string     `gorm:"type:uuid;primaryKey" json:"maliwan_data_id"`
	Odata_etag                 *string    `json:"@odata.etag"`
	No                         *string    `json:"No"`
	No_2                       *string    `json:"No_2"`
	Description                *string    `json:"Description"`
	Description_2              *string    `json:"Description_2"`
	Description_ENG            *string    `json:"Description_ENG"`
	Inventory                  *float64   `json:"Inventory"`
	Maximum_Inventory          *int32     `json:"Maximum_Inventory"`
	Safety_Stock_Quantity      *int32     `json:"Safety_Stock_Quantity"`
	Product_Group_Code         *string    `json:"Product_Group_Code"`
	Created_From_Nonstock_Item *bool      `json:"Created_From_Nonstock_Item"`
	Line_Type                  *string    `json:"Line_Type"`
	Type_Insurance             *bool      `json:"Type_Insurance"`
	Type_Registration_fee      *bool      `json:"Type_Registration_fee"`
	Type_Accessory             *bool      `json:"Type_Accessory"`
	Type_Deposit               *bool      `json:"Type_Deposit"`
	Stockkeeping_Unit_Exists   *bool      `json:"Stockkeeping_Unit_Exists"`
	Assembly_BOM               *bool      `json:"Assembly_BOM"`
	Base_Unit_of_Measure       *string    `json:"Base_Unit_of_Measure"`
	Shelf_No                   *string    `json:"Shelf_No"`
	Gen_Prod_Posting_Group     *string    `json:"Gen_Prod_Posting_Group"`
	Inventory_Posting_Group    *string    `json:"Inventory_Posting_Group"`
	Item_Category_Code         *string    `json:"Item_Category_Code"`
	VAT_Prod_Posting_Group     *string    `json:"VAT_Prod_Posting_Group"`
	Item_Disc_Group            *string    `json:"Item_Disc_Group"`
	Vendor_No                  *string    `json:"Vendor_No"`
	Vendor_Item_No             *string    `json:"Vendor_Item_No"`
	Tariff_No                  *string    `json:"Tariff_No"`
	Search_Description         *string    `json:"Search_Description"`
	Overhead_Rate              *int32     `json:"Overhead_Rate"`
	Blocked                    *bool      `json:"Blocked"`
	Last_Date_Modified         *string    `json:"Last_Date_Modified"`
	Sales_Unit_of_Measure      *string    `json:"Sales_Unit_of_Measure"`
	Replenishment_System       *string    `json:"Replenishment_System"`
	Purch_Unit_of_Measure      *string    `json:"Purch_Unit_of_Measure"`
	Lead_Time_Calculation      *string    `json:"Lead_Time_Calculation"`
	Manufacturing_Policy       *string    `json:"Manufacturing_Policy"`
	Flushing_Method            *string    `json:"Flushing_Method"`
	Assembly_Policy            *string    `json:"Assembly_Policy"`
	Item_Tracking_Code         *string    `json:"Item_Tracking_Code"`
	Promotion_Model_Code       *string    `json:"Promotion_Model_Code"`
	Default_Location_Code      *string    `json:"Default_Location_Code"`
	Qty_on_Purch_Order         *int32     `json:"Qty_on_Purch_Order"`
	Qty_on_Service_Order       *float64   `json:"Qty_on_Service_Order"`
	Qty_on_Sales_Order         *float64   `json:"Qty_on_Sales_Order"`
	Unit_Price_Price           *float64   `json:"Unit_Price_Price"`
	Brand                      *string    `json:"Brand"`
	Costing_Method             *string    `json:"Costing_Method"`
	Stockout_Warning           *string    `json:"Stockout_Warning"`
	Global_Dimension_1_Filter  *string    `json:"Global_Dimension_1_Filter"`
	Global_Dimension_2_Filter  *string    `json:"Global_Dimension_2_Filter"`
	Location_Filter            *string    `json:"Location_Filter"`
	Drop_Shipment_Filter       *string    `json:"Drop_Shipment_Filter"`
	Variant_Filter             *string    `json:"Variant_Filter"`
	Lot_No_Filter              *string    `json:"Lot_No_Filter"`
	Serial_No_Filter           *string    `json:"Serial_No_Filter"`
	Date_Filter                *string    `json:"Date_Filter"`
	Branch                     *string    `gorm:"index" json:"Branch"`
	CreatedAt                  *time.Time `json:"created_at"`
}

type Maliwan_Item_Category_Code struct {
	Item_Category_CodeID string     `gorm:"type:uuid;primaryKey" json:"item_category_code_id"`
	Item_Category_Code   *string    `json:"Item_Category_Code"`
	CreatedAt            *time.Time `json:"created_at"`
}
