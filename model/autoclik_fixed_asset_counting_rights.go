package model

type Autoclik_Fixed_Asset_Counting_Rights struct {
	Autoclik_Fixed_Asset_Counting_RightsID string                     `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_counting_rights_id"`
	UserID                                 string                     `json:"user_id"`
	User                                   User                       `validate:"-"`
	Autoclik_Fixed_Asset_CountID           string                     `json:"autoclik_fixed_asset_count_id"`
	Autoclik_Fixed_Asset_Count             Autoclik_Fixed_Asset_Count `validate:"-"`
}
