package model

type Maliwan_Counting_Trigger struct {
	Maliwan_Counting_TriggerID string        `gorm:"type:uuid;primaryKey" json:"maliwan_counting_trigger_id"`
	Working_time               string        `gorm:"working_time"`
	Day_time                   string        `gorm:"day_time"`
	Maliwan_countID            string        `json:"maliwan_count_id"`
	Maliwan_count              Maliwan_count `validate:"-"`
}
