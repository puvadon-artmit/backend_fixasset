package model

import (
	"time"

	"gorm.io/gorm"
)

type Main_branch struct {
	MainBranchID                string         `gorm:"type:uuid;primaryKey" json:"main_branch_id"`
	MainBranchName              *string        `json:"main_branch_name"`
	Company_name                *string        `json:"company_name"`
	Company_branch_name         *string        `json:"company_branch_name"`
	Company_branch_name_en      *string        `json:"company_branch_name_en"`
	Company_branch_no           *string        `json:"company_branch_no"`
	Company_address             *string        `json:"company_address"`
	Company_taxid               *string        `json:"company_taxid"`
	Company_header              *string        `json:"company_header"`
	BranchRef1                  *string        `json:"branchRef1"`
	BranchRef2                  *string        `json:"branchRef2"`
	TokenLineCreditNotification *string        `json:"tokenLineCreditNotification"`
	Navurl                      *string        `json:"navurl"`
	Parent                      int64          `json:"parent"`
	CreatedAt                   *time.Time     `json:"created_at"`
	UpdatedAt                   *time.Time     `json:"updated_at"`
	DeletedAt                   gorm.DeletedAt `gorm:"index"`
}

type Main_Branch_Story struct {
	MainBranchStoryID      string         `gorm:"type:uuid;primaryKey" json:"main_branch_story_id"`
	MainBranchStoryName    string         `json:"main_branch_story_name"`
	MainBranchStoryDetails string         `json:"main_branch_story_details"`
	CreatedAt              *time.Time     `json:"created_at"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	UserID                 string         `json:"user_id"`
	User                   User           `validate:"-"`
	MainBranchID           string         `json:"main_branch_id"`
	MainBranch             Main_branch    `validate:"-"`
}
