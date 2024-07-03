package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_Update_Story struct {
	Autoclik_Update_StoryID string         `gorm:"type:uuid;primaryKey" json:"autoclik_update_story_id"`
	Autoclik_Update_Name    string         `grom:"default:''" json:"autoclik_update_name"`
	Group_api               string         `grom:"default:''" json:"group_api"`
	CreatedAt               *time.Time     `json:"created_at"`
	UpdatedAt               *time.Time     `json:"updated_at"`
	DeletedAt               gorm.DeletedAt `gorm:"index"`
	UserID                  string         `json:"user_id"`
	User                    User           `validate:"-"`
}
