package model

import (
	"time"

	"gorm.io/gorm"
)

type Assets struct {
	AssetsID       string `gorm:"type:uuid;primaryKey" json:"assets_id"`
	Model_name     string `grom:"default:''" json:"model_name"`
	Manufacturer   string `grom:"default:''" json:"manufacturer"`
	Serial_Code    string `grom:"default:''" json:"serial_code"`
	Type           string `grom:"default:''" json:"type"`
	Model          string `grom:"default:''" json:"model"`
	Branch         string `grom:"default:''" json:"branch"`
	Username       string `grom:"default:''" json:"username"`
	Property_code  string `grom:"default:''" json:"property_code"`
	Status         string `grom:"default:''" json:"status"`
	Group_hardware string `grom:"default:''" json:"group_hardware"`
	Group          string `grom:"default:''" json:"group"`
	User_hardware  string `grom:"default:''" json:"user_hardware"`
	Phone_number   string `grom:"default:''" json:"phone_number"`
	Posting_group  string `grom:"default:''" json:"posting_group"`
	// Latest_time    time.Time `json:"latest_time"`
	Latest_time   string         `grom:"default:''" json:"latest_time"`
	ResponsibleID string         `json:"responsible_id"`
	Responsible   Responsible    `validate:"-"`
	Comment1      string         `grom:"default:''" json:"comment1"`
	Comment2      string         `grom:"default:''" json:"comment2"`
	Comment3      string         `grom:"default:''" json:"comment3"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string         `json:"user_id"`
	User          User           `validate:"-"`
	CategoryID    string         `json:"category_id"`
	Category      Category       `validate:"-"`
	ItemModelID   *string        `json:"item_model_id"`
	Item_model    Item_model     `validate:"-"`
	GroundID      string         `json:"ground_id"`
	Ground        Ground         `validate:"-"`
}

type Assets_Story struct {
	Assets_storyID   string         `gorm:"type:uuid;primaryKey" json:"assets_story_id"`
	Assets_StoryName *string        `json:"assets_story_name"`
	Property_code    string         `json:"property_code"`
	Assets_Details   *string        `json:"assets_details"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
	User             User           `validate:"-"`
}
