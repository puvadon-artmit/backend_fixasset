package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_count struct {
	Autoclik_countID       string          `gorm:"type:uuid;primaryKey" json:"autoclik_count_id"`
	Plan_Code              string          `grom:"default:''" json:"plan_code"`
	Plan_Name              *string         `grom:"default:''" json:"plan_name"`
	TypeplanName           *string         `grom:"default:''" json:"typeplan_name"`
	Project_name           *string         `grom:"default:''" json:"project_name"`
	Comment                *string         `grom:"default:''" json:"comment"`
	Plan_start             *string         `grom:"default:''" json:"plan_start"`
	Plan_end               *string         `grom:"default:''" json:"plan_end"`
	Gen_Prod_Posting_Group string          `json:"gen_prod_posting_group"`
	Status                 *string         `grom:"default:''" json:"status"`
	StatusPulldata         *string         `grom:"default:''" json:"status_pull_data"`
	CreatedAt              *time.Time      `json:"created_at"`
	UpdatedAt              *time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt  `gorm:"index"`
	UserID                 string          `json:"user_id"`
	User                   User            `validate:"-"`
	Group                  string          `grom:"default:''" json:"group"`
	BranchAutoclik_ID      string          `json:"branch_autoclik_id"`
	Branch_Autoclik        Branch_Autoclik `validate:"-"`
}

type Autoclik_count_Story struct {
	Autoclik_count_StoryID string         `gorm:"type:uuid;primaryKey" json:"autoclik_count_story_id"`
	Autoclik_count_Name    string         `grom:"default:''" json:"autoclik_count_name"`
	Autoclik_count_details string         `grom:"default:''" json:"autoclik_count_details"`
	CreatedAt              *time.Time     `json:"created_at"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	UserID                 string         `json:"user_id"`
	User                   User           `validate:"-"`
	Autoclik_countID       string         `json:"autoclik_count_id"`
	Autoclik_count         Autoclik_count `validate:"-"`
}

type Autoclik_Count_Product_Group struct {
	Autoclik_Count_Product_GroupID string         `gorm:"type:uuid;primaryKey" json:"autoclik_count_product_group_id"`
	Name_Product_Group             *string        `json:"name_product_group"`
	Autoclik_countID               string         `json:"autoclik_count_id"`
	Autoclik_count                 Autoclik_count `validate:"-"`
}
