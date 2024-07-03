package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Asset_check struct {
	Asset_checkID      string         `gorm:"type:uuid;primaryKey" json:"asset_check_id"`
	Additional_notes   *string        `grom:"default:''" json:"additional_notes"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Property_code      *string        `grom:"default:''" json:"property_code"`
	StatusAsset        *string        `grom:"default:''" json:"statusasset"`
	Quantity_remaining int64          `json:"quantity_remaining"`
	Round_CountID      uuid.UUID      `gorm:"type:uuid" json:"round_count_id"`
	Round_Count        Round_Count    `validate:"-"`
	Photos_check       []Photos_check `gorm:"many2many:photocheck;" json:"photos_check"`
	UserID             string         `json:"user_id"`
	User               User           `validate:"-"`
}

type PhotoCheck struct {
	gorm.Model
	Asset_checkID  string `gorm:"type:uuid"`
	Photos_checkID string `gorm:"type:uuid"`
}

type Photos_check struct {
	Photos_checkID string `gorm:"type:uuid;primaryKey" json:"photos_check_id"`
	Photos         string `json:"photos"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      gorm.DeletedAt
}
