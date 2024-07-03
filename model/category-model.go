package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CategoryID       string         `gorm:"type:uuid;primaryKey" json:"category_id"`
	CategoryName     string         `grom:"default:''" json:"category_name"`
	CategoryImage    string         `grom:"default:''" json:"category_image"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
	User             User
	Main_Category_ID string        `json:"main_category_id"`
	Main_Category    Main_Category `validate:"-"`
}

type Category_Story struct {
	Category_Story_ID    string         `gorm:"type:uuid;primaryKey" json:"category_story_id"`
	Category_Story_Name  string         `grom:"default:''" json:"category_story_name"`
	Category_Story_Dtail string         `grom:"default:''" json:"category_story_detail"`
	CreatedAt            *time.Time     `json:"created_at"`
	UpdatedAt            *time.Time     `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	UserID               string         `json:"user_id"`
	User                 User           `validate:"-"`
	CategoryID           string         `json:"category_id"`
	Category             Category       `validate:"-"`
}
