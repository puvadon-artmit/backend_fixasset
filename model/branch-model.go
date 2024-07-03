package model

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	BranchID     string         `gorm:"type:uuid;primaryKey" json:"id_branch"`
	ZipCode      string         `grom:"default:''" json:"zip_code"`
	County       string         `grom:"default:''" json:"county"`
	Province     string         `grom:"default:''" json:"province"`
	Comment      string         `grom:"default:''" json:"comment"`
	Building     string         `grom:"default:''" json:"building"`
	Address      string         `grom:"default:''" json:"address"`
	Town         string         `grom:"default:''" json:"town"`
	RoomNumber   string         `grom:"default:''" json:"room_number"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       string         `json:"user_id"`
	GroupID      string         `json:"group_id"`
	Group        Group          `validate:"-"`
	MainBranchID string         `json:"main_branch_id"`
	Main_branch  Main_branch    `validate:"-"`
}

type Branch_Story struct {
	BranchStoryID      string         `gorm:"type:uuid;primaryKey" json:"branch_story_id"`
	BranchStoryName    string         `grom:"default:''" json:"branch_story_name"`
	BranchStoryDetails string         `grom:"default:''" json:"branch_story_details"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	UserID             string         `json:"user_id"`
	User               User           `validate:"-"`
	BranchID           string         `json:"branch_id"`
	Branch             Branch         `validate:"-"`
}
