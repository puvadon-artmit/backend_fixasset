package model

import (
	"time"

	"gorm.io/gorm"
)

type Type_things struct {
	TypeID     string         `gorm:"type:uuid;primaryKey" json:"type_id"`
	TypeName   *string        `json:"type_name" validate:"required"`
	Comment    *string        `json:"comment"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserID     string         `json:"user_id"`
	CategoryID string         `json:"category_id"`
	Category   Category       `validate:"-"`
}

type Type_things_story struct {
	Type_things_storyID       string         `gorm:"type:uuid;primaryKey" json:"type_things_story_id"`
	Type_things_story_Name    string         `json:"type_things_story_name" validate:"required"`
	Type_things_story_Details *string        `json:"type_things_story_details"`
	CreatedAt                 *time.Time     `json:"created_at"`
	UpdatedAt                 *time.Time     `json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
	Type_things_ID            string         `json:"type_id"`
	Type_things               Type_things    `validate:"-"`

	UserID string `json:"user_id"`
	User   User   `validate:"-"`
}
