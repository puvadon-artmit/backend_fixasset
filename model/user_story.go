package model

import (
	"time"

	"gorm.io/gorm"
)

type User_story struct {
	User_storyID string         `gorm:"type:uuid;primaryKey" json:"user_story_id"`
	User_name    *string        `json:"user_name"`
	User_detail  *string        `json:"user_detail"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       string         `json:"user_id"`
	User         User           `validate:"-"`
}
