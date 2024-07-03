package model

import (
	"time"

	"gorm.io/gorm"
)

type Count_Autoclik struct {
	Autoclik_countID   string          `gorm:"type:uuid;primaryKey" json:"autoclik_count_id"`
	Plan_Code          string          `json:"plan_code"`
	Plan_Name          *string         `json:"plan_name"`
	TypeplanName       *string         `json:"typeplan_name"`
	Project_name       *string         `json:"project_name"`
	Comment            *string         `json:"comment"`
	Plan_start         *string         `json:"plan_start"`
	Plan_end           *string         `json:"plan_end"`
	Status             *string         `json:"status"`
	Product_Group_Code *string         `json:"Product_Group_Code"`
	CreatedAt          *time.Time      `json:"created_at"`
	UpdatedAt          *time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt  `gorm:"index"`
	UserID             string          `json:"user_id"`
	User               User            `validate:"-"`
	Group              string          `json:"group"`
	BranchAutoclik_ID  string          `json:"branch_autoclik_id"`
	Branch_Autoclik    Branch_Autoclik `validate:"-"`
}
