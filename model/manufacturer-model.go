package model

import (
	"time"

	"gorm.io/gorm"
)

type Manufacturer struct {
	ManufacturerID   string         `gorm:"type:uuid;primaryKey" json:"manufacturer_id"`
	ManufacturerName *string        `grom:"default:''" json:"manufacturer_name" validate:"required"`
	Comment          *string        `grom:"default:''" json:"comment"`
	CreatedAt        *time.Time     `json:"created_at"`
	UpdatedAt        *time.Time     `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type Manufacturer_Story struct {
	Manufacturer_StoryID       string         `gorm:"type:uuid;primaryKey" json:"manufacturer_story_id"`
	Manufacturer_Story_Name    *string        `grom:"default:''" json:"manufacturer_story_name"`
	Manufacturer_Story_Details *string        `grom:"default:''" json:"manufacturer_story_details"`
	CreatedAt                  *time.Time     `json:"created_at"`
	UpdatedAt                  *time.Time     `json:"updated_at"`
	DeletedAt                  gorm.DeletedAt `gorm:"index"`
	UserID                     string         `json:"user_id"`
	User                       User           `validate:"-"`
	ManufacturerID             string         `json:"manufacturer_id"`
	Manufacturer               Manufacturer   `validate:"-"`
}
