package model

import (
	"time"

	"gorm.io/gorm"
)

type Branch_Autoclik struct {
	BranchAutoclik_ID string         `gorm:"type:uuid;primaryKey" json:"branchautoclik_id"`
	Branch_Name       string         `json:"branch_name"`
	Branch_Code       string         `json:"branch_code"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
