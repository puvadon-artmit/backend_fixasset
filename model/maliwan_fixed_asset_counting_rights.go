package model

type Maliwan_Fixed_Asset_Counting_Rights struct {
	Maliwan_Fixed_Asset_Counting_RightsID string                    `gorm:"type:uuid;primaryKey" json:"maliwan_fixed_asset_counting_rights_id"`
	UserID                                string                    `json:"user_id"`
	User                                  User                      `validate:"-"`
	Maliwan_Fixed_Asset_CountID           string                    `json:"maliwan_fixed_asset_count_id"`
	Maliwan_Fixed_Asset_Count             Maliwan_Fixed_Asset_Count `validate:"-"`
}
