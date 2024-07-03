package model

import (
	"time"

	"gorm.io/gorm"
)

type Maliwan_Fixed_Asset_Round_Count struct {
	Maliwan_Fixed_Asset_Round_CountID string                    `gorm:"type:uuid;primaryKey" json:"maliwan_fixed_asset_round_count_id"`
	Round                             *string                   `json:"round"`
	Status                            *string                   `json:"status"`
	CreatedAt                         *time.Time                `json:"created_at"`
	UpdatedAt                         *time.Time                `json:"updated_at"`
	DeletedAt                         gorm.DeletedAt            `gorm:"index"`
	Maliwan_Fixed_Asset_CountID       string                    `json:"maliwan_fixed_asset_count_id"`
	Maliwan_Fixed_Asset_Count         Maliwan_Fixed_Asset_Count `validate:"-"`
}

type Maliwan_Fixed_Asset_Round_Count_Story struct {
	Maliwan_Fixed_Asset_Round_Count_StoryID string                          `gorm:"type:uuid;primaryKey" json:"maliwan_fixed_asset_round_count_story_id"`
	Maliwan_Round_Name_Strory               *string                         `json:"maliwan_round_name_strory"`
	Maliwan_Round_Strory_Detail             *string                         `json:"maliwan_round_detail_strory"`
	CreatedAt                               *time.Time                      `json:"created_at"`
	UpdatedAt                               *time.Time                      `json:"updated_at"`
	DeletedAt                               gorm.DeletedAt                  `gorm:"index"`
	Maliwan_Fixed_Asset_Round_CountID       string                          `json:"maliwan_fixed_asset_round_count_id"`
	Maliwan_Fixed_Asset_Round_Count         Maliwan_Fixed_Asset_Round_Count `validate:"-"`
	UserID                                  string                          `json:"user_id"`
	User                                    User                            `validate:"-"`
}
