package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_Fixed_Asset_Round_Count struct {
	Autoclik_Fixed_Asset_Round_CountID string                     `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_round_count_id"`
	Round                              *string                    `json:"round"`
	Status                             *string                    `json:"status"`
	CreatedAt                          *time.Time                 `json:"created_at"`
	UpdatedAt                          *time.Time                 `json:"updated_at"`
	DeletedAt                          gorm.DeletedAt             `gorm:"index"`
	Autoclik_Fixed_Asset_CountID       string                     `json:"autoclik_fixed_asset_count_id"`
	Autoclik_Fixed_Asset_Count         Autoclik_Fixed_Asset_Count `validate:"-"`
}

type Autoclik_Fixed_Asset_Round_Count_Story struct {
	Autoclik_Fixed_Asset_Round_Count_StoryID string                           `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_round_count_story_id"`
	Autoclik_Round_Name_Strory               *string                          `json:"autoclik_round_name_strory"`
	Autoclik_Round_Strory_Detail             *string                          `json:"autoclik_round_detail_strory"`
	CreatedAt                                *time.Time                       `json:"created_at"`
	UpdatedAt                                *time.Time                       `json:"updated_at"`
	DeletedAt                                gorm.DeletedAt                   `gorm:"index"`
	Autoclik_Fixed_Asset_Round_CountID       string                           `json:"autoclik_fixed_asset_round_count_id"`
	Autoclik_Fixed_Asset_Round_Count         Autoclik_Fixed_Asset_Round_Count `validate:"-"`
	UserID                                   string                           `json:"user_id"`
	User                                     User                             `validate:"-"`
}
