package model

type Itemabinmaliwanodata struct {
	Odata_context string             `json:"@odata.context"`
	Value         []Item_bin_maliwan `json:"value"`
}

type Item_bin_maliwan struct {
	ItemabinmaliwanID    string   `gorm:"type:uuid;primaryKey" json:"itema_bin_maliwan_id"`
	Location_Code        *string  `json:"Location_Code"`
	Bin_Code             *string  `json:"bin_code"`
	Item_No              *string  `json:"Item_No"`
	Variant_Code         *string  `json:"Variant_Code"`
	Unit_of_Measure_Code *string  `json:"Unit_of_Measure_Code"`
	Fixed                *bool    `json:"Fixed"`
	Default              *bool    `json:"Default"`
	Dedicated            *bool    `json:"Dedicated"`
	CalcQtyUOM           *float64 `json:"CalcQtyUOM"`
	Quantity_Base        *int64   `json:"Quantity_Base"`
	Bin_Type_Code        *string  `json:"Bin_Type_Code"`
	Zone_Code            *string  `json:"Zone_Code"`
	Lot_No_Filter        *string  `json:"Lot_No_Filter"`
	Serial_No_Filter     *string  `json:"Serial_No_Filter"`
	Branch               *string  `json:"Branch"`
}
