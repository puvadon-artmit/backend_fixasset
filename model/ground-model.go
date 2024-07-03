package model

import (
	"time"

	"gorm.io/gorm"
)

type Ground struct {
	GroundID      string         `gorm:"type:uuid;primaryKey" json:"ground_id"`
	GroundName    *string        `grom:"default:''" json:"ground_name"`
	Location_code *string        `grom:"default:''" json:"location_code"`
	Building      *string        `grom:"default:''" json:"building"`
	Floor         *string        `grom:"default:''" json:"floor"`
	Room          *string        `grom:"default:''" json:"room"`
	GroundImage   *string        `grom:"default:''" json:"ground_image"`
	Comment       *string        `grom:"default:''" json:"comment"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Ground_Story struct {
	Ground_storyID   string         `gorm:"type:uuid;primaryKey" json:"ground_story_id"`
	Ground_StoryName *string        `grom:"default:''" json:"groundstory_name"`
	Ground_Details   *string        `grom:"default:''" json:"ground_details"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           string         `json:"user_id"`
	User             User           `validate:"-"`
	GroundID         string         `json:"ground_id"`
	Ground           Ground         `validate:"-"`
}
