package model

import (
	"time"

	"gorm.io/gorm"
)

type Maliwan_count struct {
	Maliwan_countID        string         `gorm:"type:uuid;primaryKey" json:"maliwan_count_id"`
	Plan_Code              string         `grom:"default:''" json:"plan_code"`
	Plan_Name              *string        `grom:"default:''" json:"plan_name"`
	TypeplanName           *string        `grom:"default:''" json:"typeplan_name"`
	Project_name           *string        `grom:"default:''" json:"project_name"`
	Comment                *string        `grom:"default:''" json:"comment"`
	Plan_start             *string        `grom:"default:''" json:"plan_start"`
	Plan_end               *string        `grom:"default:''" json:"plan_end"`
	Status                 *string        `grom:"default:''" json:"status"`
	StatusPulldata         *string        `grom:"default:''" json:"status_pull_data"`
	CreatedAt              *time.Time     `json:"created_at"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	UserID                 string         `json:"user_id"`
	User                   User           `validate:"-"`
	Branch_Maliwan         *string        `grom:"default:''" json:"branch_maliwan"`
	Gen_Prod_Posting_Group string         `json:"gen_prod_posting_group"`
}

type Maliwan_counts_story struct {
	Maliwan_counts_storyID      string         `gorm:"type:uuid;primaryKey" json:"maliwan_counts_story_id"`
	Maliwan_count_story_name    string         `grom:"default:''" json:"maliwan_count_story_name"`
	Maliwan_counts_story_detail string         `grom:"default:''" json:"maliwan_counts_story_detail"`
	CreatedAt                   *time.Time     `json:"created_at"`
	UpdatedAt                   *time.Time     `json:"updated_at"`
	DeletedAt                   gorm.DeletedAt `gorm:"index"`
	UserID                      string         `json:"user_id"`
	User                        User           `validate:"-"`
	Maliwan_countID             string         `json:"maliwan_counts_id"`
	Maliwan_count               Maliwan_count  `validate:"-"`
}

type Maliwan_Counts_Item_Category_Code struct {
	Maliwan_Counts_Item_Category_CodeID string        `gorm:"type:uuid;primaryKey" json:"maliwan_counts_item_category_code_id"`
	Item_Category_Code                  *string       `json:"item_category_code"`
	Maliwan_countID                     string        `json:"maliwan_count_id"`
	Maliwan_count                       Maliwan_count `validate:"-"`
}
