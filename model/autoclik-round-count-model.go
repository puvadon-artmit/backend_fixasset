package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_Round_Count struct {
	Autoclik_Round_CountID string         `gorm:"type:uuid;primaryKey" json:"autoclik_round_count_id"`
	Round                  *string        `json:"round"`
	Status                 *string        `json:"status"`
	CreatedAt              *time.Time     `json:"created_at"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	Autoclik_countID       string         `json:"autoclik_count_id"`
	Autoclik_count         Autoclik_count `validate:"-"`
}

type Autoclik_Round_Count_Story struct {
	Autoclik_Round_Count_StoryID string               `gorm:"type:uuid;primaryKey" json:"autoclik_round_count_story_id"`
	Autoclik_Round_Name_Strory   *string              `json:"autoclik_round_name_strory"`
	Autoclik_Round_Strory_Detail *string              `json:"autoclik_round_detail_strory"`
	CreatedAt                    *time.Time           `json:"created_at"`
	UpdatedAt                    *time.Time           `json:"updated_at"`
	DeletedAt                    gorm.DeletedAt       `gorm:"index"`
	Autoclik_Round_CountID       string               `json:"autoclik_round_count_id"`
	Autoclik_Round_Count         Autoclik_Round_Count `validate:"-"`
	UserID                       string               `json:"user_id"`
	User                         User                 `validate:"-"`
}
