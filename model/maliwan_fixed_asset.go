package model

type Item_Maliwan_Fixed_Asset struct {
	Odata_context string                `json:"@odata.context"`
	Value         []Maliwan_Fixed_Asset `json:"value"`
}

type Maliwan_Fixed_Asset struct {
	Maliwan_Fixed_AssetId   string  `gorm:"type:uuid;primaryKey" json:"maliwan_fixed_asset_id"`
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
	Branch                  string  `grom:"default:''" json:"branch"`
}
