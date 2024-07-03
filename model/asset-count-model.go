package model

import (
	"time"

	"gorm.io/gorm"
)

type Asset_count struct {
	Asset_countID  string         `gorm:"type:uuid;primaryKey" json:"asset_count_id"`
	Plan_Code      string         `grom:"default:''" json:"plan_code"`
	Plan_Name      *string        `grom:"default:''" json:"plan_name"`
	TypeplanName   *string        `grom:"default:''" json:"typeplan_name"`
	Project_name   *string        `grom:"default:''" json:"project_name"`
	Comment        *string        `grom:"default:''" json:"comment"`
	Plan_start     *string        `grom:"default:''" json:"plan_start"`
	Plan_end       *string        `grom:"default:''" json:"plan_end"`
	Status         *string        `grom:"default:''" json:"status"`
	StatusPulldata *string        `grom:"default:''" json:"status_pull_data"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserID         string         `json:"user_id"`
	User           User           `validate:"-"`
	Group          string         `json:"group"`
	MainBranchID   string         `json:"main_branch_id"`
	Main_branch    Main_branch    `validate:"-"`
	// Company       string         `json:"company"`
	// Main_Category_ID *string        `json:"main_category_id"`
	// CategoryID       *string        `json:"category_id"`
}

type Asset_count_story struct {
	Asset_count_storyID string         `gorm:"type:uuid;primaryKey" json:"asset_count_story_id"`
	Asset_count_Name    string         `json:"asset_count_name"`
	Asset_count_details string         `json:"asset_count_details"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	UserID              string         `json:"user_id"`
	User                User           `validate:"-"`
	Asset_countID       string         `json:"asset_count_id"`
	Asset_count         Asset_count    `validate:"-"`
}

type Asset_count_Main_Category struct {
	Asset_count_Main_CategoryID string        `gorm:"type:uuid;primaryKey" json:"asset_count_main_category_id"`
	Main_Category_ID            *string       `json:"main_category_id"`
	Main_Category               Main_Category `validate:"-"`
	Asset_countID               string        `json:"asset_count_id"`
	Asset_count                 Asset_count   `validate:"-"`
}

type Asset_count_Category struct {
	Asset_count_CategoryID string      `gorm:"type:uuid;primaryKey" json:"asset_count_category_id"`
	Category_ID            *string     `json:"category_id"`
	Category               Category    `validate:"-"`
	Asset_countID          string      `json:"asset_count_id"`
	Asset_count            Asset_count `validate:"-"`
}

type Inspection_story struct {
	Inspection_storyID  string         `gorm:"type:uuid;primaryKey" json:"inspection_story_id"`
	Inspection_name     *string        `json:"inspection_name"`
	Inspection_Username *string        `json:"inspection_username"`
	Inspection_Details  *string        `json:"inspection_details"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	Asset_countID       string         `json:"asset_count_id"`
	Asset_count         Asset_count    `validate:"-"`
}
