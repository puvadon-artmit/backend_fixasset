package model

type Maliwan_Counting_Rights struct {
	Maliwan_Counting_RightsID string        `gorm:"type:uuid;primaryKey" json:"maliwan_counting_rights_id"`
	UserID                    string        `json:"user_id"`
	User                      User          `validate:"-"`
	Maliwan_countID           string        `json:"maliwan_count_id"`
	Maliwan_count             Maliwan_count `validate:"-"`
}
