package model

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	GroupID   string         `gorm:"type:uuid;primaryKey" json:"group_id"`
	GroupName *string        `grom:"default:''" json:"name_group"`
	Comment   *string        `grom:"default:''" json:"comment"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string         `json:"user_id"`
}

type Group_Story struct {
	GroupStoryID      string         `gorm:"type:uuid;primaryKey" json:"group_story_id"`
	GroupStoryName    string         `grom:"default:''" json:"group_story_name"`
	GroupStoryDetails *string        `grom:"default:''" json:"group_story_details"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	UserID            string         `json:"user_id"`
	User              User           `validate:"-"`
	GroupID           string         `json:"group_id"`
	Group             Group          `validate:"-"`
}
