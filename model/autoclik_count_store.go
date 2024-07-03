package model

type Autoclik_Count_Store struct {
	Autoclik_Count_StoreID string         `gorm:"type:uuid;primaryKey" json:"autoclik_count_store_id"`
	Odata_etag             *string        `grom:"default:''" json:"@odata.etag"`
	Location_Code          *string        `grom:"default:''" json:"Location_Code"`
	Bin_Code               *string        `grom:"default:''" json:"Bin_Code"`
	Serial_Code            string         `grom:"default:''" json:"serial_code"`
	Item_No                *string        `grom:"default:''" json:"Item_No"`
	Variant_Code           *string        `grom:"default:''" json:"Variant_Code"`
	Unit_of_Measure_Code   *string        `grom:"default:''" json:"Unit_of_Measure_Code"`
	ItemDesc               *string        `grom:"default:''" json:"ItemDesc"`
	ItemDesc2              *string        `grom:"default:''" json:"ItemDesc2"`
	ItemCategory           *string        `grom:"default:''" json:"ItemCategory"`
	ProductGroup           *string        `grom:"default:''" json:"ProductGroup"`
	Fixed                  bool           `grom:"default:''" json:"Fixed"`
	Default                bool           `grom:"default:''" json:"Default"`
	Dedicated              bool           `grom:"default:''" json:"Dedicated"`
	CalcQtyUOM             float64        `grom:"default:''" json:"CalcQtyUOM"`
	Quantity_Base          int64          `grom:"default:''" json:"Quantity_Base"`
	Bin_Type_Code          *string        `grom:"default:''" json:"Bin_Type_Code"`
	Zone_Code              *string        `grom:"default:''" json:"Zone_Code"`
	Lot_No_Filter          *string        `grom:"default:''" json:"Lot_No_Filter"`
	Serial_No_Filter       *string        `grom:"default:''" json:"Serial_No_Filter"`
	Autoclik_countID       string         `json:"autoclik_count_id"`
	Autoclik_count         Autoclik_count `validate:"-"`
}
