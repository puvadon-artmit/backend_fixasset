package model

type Autoclik_Counting_Trigger struct {
	Autoclik_Counting_TriggerID string         `gorm:"type:uuid;primaryKey" json:"autoclik_counting_trigger_id"`
	Working_time                string         `gorm:"working_time"`
	Day_time                    string         `gorm:"day_time"`
	Autoclik_countID            string         `json:"autoclik_count_id"`
	Autoclik_count              Autoclik_count `validate:"-"`
}
