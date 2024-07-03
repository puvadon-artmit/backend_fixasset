package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Maliwan_check struct {
	Maliwan_checkID       string                 `gorm:"type:uuid;primaryKey" json:"maliwan_check_id"`
	Additional_notes      *string                `grom:"default:''" json:"additional_notes"`
	CreatedAt             *time.Time             `json:"created_at"`
	UpdatedAt             *time.Time             `json:"updated_at"`
	DeletedAt             gorm.DeletedAt         `gorm:"index"`
	StatusMaliwan         *string                `grom:"default:''" json:"statusmaliwan"`
	No                    *string                `grom:"default:''" json:"No"`
	Bin_Code              *string                `grom:"default:''" json:"bin_code"`
	Location_Code         *string                `grom:"default:''" json:"location_code"`
	Quantity_Remaining    int64                  `grom:"default:''" json:"quantity_remaining"`
	Maliwan_Round_CountID uuid.UUID              `gorm:"type:uuid" json:"maliwan_round_count_id"`
	Maliwan_Round_Count   Maliwan_Round_Count    `validate:"-"`
	Maliwan_Photos_check  []Maliwan_Photos_check `gorm:"many2many:photomaliwancheck;" json:"maliwan_photos_check"`
	UserID                string                 `json:"user_id"`
	User                  User                   `validate:"-"`
}

type PhotoMaliwanCheck struct {
	gorm.Model
	Maliwan_checkID        string `gorm:"type:uuid"`
	Maliwan_Photos_checkID string `gorm:"type:uuid"`
}

type Maliwan_Photos_check struct {
	Maliwan_Photos_checkID string `gorm:"type:uuid;primaryKey" json:"maliwan_photos_check_id"`
	Photos                 string `json:"photos"`
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
	DeletedAt              gorm.DeletedAt
}

type Maliwan_check_Story struct {
	Maliwan_check_StoryID string              `gorm:"type:uuid;primaryKey" json:"maliwan_check_story_id"`
	Maliwan_check_Name    string              `grom:"default:''" json:"maliwan_check_story_name"`
	Maliwan_check_details string              `grom:"default:''" json:"maliwan_check_story_details"`
	No                    *string             `json:"No"`
	CreatedAt             *time.Time          `json:"created_at"`
	UpdatedAt             *time.Time          `json:"updated_at"`
	DeletedAt             gorm.DeletedAt      `gorm:"index"`
	UserID                string              `json:"user_id"`
	User                  User                `validate:"-"`
	Maliwan_Round_CountID string              `json:"maliwan_round_count_id"`
	Maliwan_Round_Count   Maliwan_Round_Count `validate:"-"`
}
