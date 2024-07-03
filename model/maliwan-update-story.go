package model

import (
	"time"

	"gorm.io/gorm"
)

type Maliwan_Update_Story struct {
	Maliwan_Update_StoryID string         `gorm:"type:uuid;primaryKey" json:"maliwan_update_story_id"`
	Maliwan_Update_Name    string         `grom:"default:''" json:"maliwan_update_name"`
	Group_api              string         `grom:"default:''" json:"group_api"`
	CreatedAt              *time.Time     `json:"created_at"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	UserID                 string         `json:"user_id"`
	User                   User           `validate:"-"`
}

type Maliwan_Update_Fixed_Asset_Story struct {
	Maliwan_Update_Fixed_Asset_StoryID string         `gorm:"type:uuid;primaryKey" json:"maliwan_update_fixed_asset_story_id"`
	Maliwan_FA_Update_Name             string         `grom:"default:''" json:"maliwan_fa_update_name"`
	CreatedAt                          *time.Time     `json:"created_at"`
	UpdatedAt                          *time.Time     `json:"updated_at"`
	DeletedAt                          gorm.DeletedAt `gorm:"index"`
	UserID                             string         `json:"user_id"`
	User                               User           `validate:"-"`
}
