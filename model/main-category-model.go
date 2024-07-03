package model

import (
	"time"

	"gorm.io/gorm"
)

type Main_Category struct {
	Main_Category_ID    string         `gorm:"type:uuid;primaryKey" json:"main_category_id"`
	Code_ID             string         `json:"code_id"`
	Main_name           string         `grom:"default:''" json:"main_name"`
	Main_Category_Photo string         `json:"main_category_photo"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type Main_Category_story struct {
	Main_Category_storyID   string         `gorm:"type:uuid;primaryKey" json:"main_category_story_id"`
	Main_Category_StoryName *string        `grom:"default:''" json:"main_category_story_name"`
	Main_Category_Details   *string        `grom:"default:''" json:"main_category_details"`
	CreatedAt               *time.Time     `json:"created_at"`
	UpdatedAt               *time.Time     `json:"updated_at"`
	DeletedAt               gorm.DeletedAt `gorm:"index"`
	Main_Category_ID        string         `json:"main_category_id"`
	Main_Category           Main_Category  `validate:"-"`
	UserID                  string         `json:"user_id"`
	User                    User           `validate:"-"`
}
