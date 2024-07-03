package model

type Counting_rights struct {
	Counting_rightsID string      `json:"counting_rights_id"`
	UserID            string      `json:"user_id"`
	User              User        `validate:"-"`
	Asset_countID     string      `json:"asset_count_id"`
	Asset_count       Asset_count `validate:"-"`
}
