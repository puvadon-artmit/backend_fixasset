package model

type Autoclik_Counting_Rights struct {
	Autoclik_Counting_RightsID string         `gorm:"type:uuid;primaryKey" json:"autoclik_counting_rights_id"`
	UserID                     string         `json:"user_id"`
	User                       User           `validate:"-"`
	Autoclik_countID           string         `json:"autoclik_count_id"`
	Autoclik_count             Autoclik_count `validate:"-"`
}
