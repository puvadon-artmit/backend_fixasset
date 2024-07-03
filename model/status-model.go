package model

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	StatusID   string         `gorm:"type:uuid;primaryKey" json:"status_id"`
	StatusName *string        `json:"status_name" validate:"required"`
	Comment    *string        `json:"comment"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Status_story struct {
	Status_story_ID     string         `json:"status_story_id"`
	Status_story_Name   *string        `json:"status_story_name"`
	Status_story_Detail *string        `json:"status_story_detail"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	UserID              string         `json:"user_id"`
	User                User           `validate:"-"`
	StatusID            string         `json:"status_id"`
	Status              Status         `validate:"-"`
}
