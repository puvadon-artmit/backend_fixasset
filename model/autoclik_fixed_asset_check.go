package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_Fixed_Asset_Check struct {
	Autoclik_Fixed_Asset_CheckID       string                           `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_check_id"`
	Additional_notes                   *string                          `grom:"default:''" json:"additional_notes"`
	CreatedAt                          *time.Time                       `json:"created_at"`
	UpdatedAt                          *time.Time                       `json:"updated_at"`
	DeletedAt                          gorm.DeletedAt                   `gorm:"index"`
	No                                 *string                          `grom:"default:''" json:"no"`
	FA_Location_Code                   *string                          `grom:"default:''" json:"fa_location_code"`
	Quantity_remaining                 int64                            `grom:"default:''" json:"quantity_remaining"`
	Status                             *string                          `json:"status"`
	UserID                             string                           `json:"user_id"`
	User                               User                             `validate:"-"`
	Autoclik_Fixed_Asset_Round_CountID string                           `json:"autoclik_fixed_asset_round_count_id"`
	Autoclik_Fixed_Asset_Round_Count   Autoclik_Fixed_Asset_Round_Count `validate:"-"`
}

type Autoclik_Fixed_Asset_Photos_check struct {
	Autoclik_Fixed_Asset_Photos_checkID string `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_photos_check_id"`
	Photos                              string `json:"photos"`
	CreatedAt                           *time.Time
	UpdatedAt                           *time.Time
	DeletedAt                           gorm.DeletedAt
}

type Autoclik_Fixed_Asset_check_Story struct {
	Autoclik_Fixed_Asset_check_StoryID string                           `gorm:"type:uuid;primaryKey" json:"autoclik_check_story_id"`
	Autoclik_Fixed_Asset_check_Name    string                           `grom:"default:''" json:"autoclik_fixed_asset_check_story_name"`
	Autoclik_Fixed_Asset_check_details string                           `grom:"default:''" json:"autoclik_fixed_asset_check_story_details"`
	No                                 *string                          `json:"no"`
	CreatedAt                          *time.Time                       `json:"created_at"`
	UpdatedAt                          *time.Time                       `json:"updated_at"`
	DeletedAt                          gorm.DeletedAt                   `gorm:"index"`
	UserID                             string                           `json:"user_id"`
	User                               User                             `validate:"-"`
	Autoclik_Fixed_Asset_Round_CountID string                           `json:"autoclik_fixed_asset_round_count_id"`
	Autoclik_Fixed_Asset_Round_Count   Autoclik_Fixed_Asset_Round_Count `validate:"-"`
}
