package model

import (
	"time"

	"gorm.io/gorm"
)

type Typeplan struct {
	TypeplanID   string         `gorm:"type:uuid;primaryKey" json:"typeplan_id"`
	TypeplanName *string        `json:"typeplan_name"`
	Comment      *string        `json:"comment"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Typeplan_story struct {
	Typeplan_storyID   string         `gorm:"type:uuid;primaryKey" json:"typeplan_story_id"`
	Typeplan_StoryName *string        `json:"typeplan_story_name"`
	Typeplan_Username  *string        `json:"typeplan_username"`
	Typeplan_Details   *string        `json:"typeplan_details"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	UserID             string         `json:"user_id"`
	TypeplanID         string         `json:"typeplan_id"`
	Typeplan           Typeplan       `validate:"-"`
}
