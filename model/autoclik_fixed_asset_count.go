package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_Fixed_Asset_Count struct {
	Autoclik_Fixed_Asset_CountID string         `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_count_id"`
	Plan_Code                    string         `grom:"default:''" json:"plan_code"`
	Plan_Name                    *string        `grom:"default:''" json:"plan_name"`
	TypeplanName                 *string        `grom:"default:''" json:"typeplan_name"`
	Project_name                 *string        `grom:"default:''" json:"project_name"`
	Comment                      *string        `grom:"default:''" json:"comment"`
	Plan_start                   *string        `grom:"default:''" json:"plan_start"`
	Plan_end                     *string        `grom:"default:''" json:"plan_end"`
	FA_Location_Code             *string        `grom:"default:''" json:"fa_location_code"`
	Status                       *string        `grom:"default:''" json:"status"`
	StatusPulldata               *string        `grom:"default:''" json:"status_pull_data"`
	CreatedAt                    *time.Time     `json:"created_at"`
	UpdatedAt                    *time.Time     `json:"updated_at"`
	DeletedAt                    gorm.DeletedAt `gorm:"index"`
	UserID                       string         `json:"user_id"`
	User                         User           `validate:"-"`
}

type Autoclik_Fixed_Asset_count_Story struct {
	Autoclik_Fixed_Asset_count_StoryID string                     `gorm:"type:uuid;primaryKey" json:"autoclik_fixed_asset_count_story_id"`
	Autoclik_Fixed_Asset_count_Name    string                     `grom:"default:''" json:"autoclik_fixed_asset_count_name"`
	Autoclik_Fixed_Asset_count_details string                     `grom:"default:''" json:"autoclik_fixed_asset_count_details"`
	CreatedAt                          *time.Time                 `json:"created_at"`
	UpdatedAt                          *time.Time                 `json:"updated_at"`
	DeletedAt                          gorm.DeletedAt             `gorm:"index"`
	UserID                             string                     `json:"user_id"`
	User                               User                       `validate:"-"`
	Autoclik_Fixed_Asset_CountID       string                     `json:"autoclik_fixed_asset_count_id"`
	Autoclik_Fixed_Asset_Count         Autoclik_Fixed_Asset_Count `validate:"-"`
}
