package model

import (
	"time"

	"gorm.io/gorm"
)

type Autoclik_check struct {
	Autoclik_checkID       string                  `gorm:"type:uuid;primaryKey" json:"autoclik_check_id"`
	Additional_notes       *string                 `grom:"default:''" json:"additional_notes"`
	CreatedAt              *time.Time              `json:"created_at"`
	UpdatedAt              *time.Time              `json:"updated_at"`
	DeletedAt              gorm.DeletedAt          `gorm:"index"`
	Item_no                *string                 `grom:"default:''" json:"item_no"`
	Quantity_remaining     int64                   `grom:"default:''" json:"quantity_remaining"`
	Bin_Code               *string                 `grom:"default:''" json:"bin_code"`
	Status                 *string                 `json:"status"`
	UserID                 string                  `json:"user_id"`
	User                   User                    `validate:"-"`
	Autoclik_Round_CountID string                  `json:"autoclik_round_count_id"`
	Autoclik_Round_Count   Autoclik_Round_Count    `validate:"-"`
	Autoclik_Photos_check  []Autoclik_Photos_check `gorm:"many2many:autoclik_allphoto;" json:"autoclik_photos_check"`
}

type Autoclik_AllPhoto struct {
	gorm.Model
	Autoclik_checkID        string `gorm:"type:uuid"`
	Autoclik_Photos_checkID string `gorm:"type:uuid"`
}

type Autoclik_Photos_check struct {
	Autoclik_Photos_checkID string `gorm:"type:uuid;primaryKey" json:"autoclik_photos_check_id"`
	Photos                  string `json:"photos"`
	CreatedAt               *time.Time
	UpdatedAt               *time.Time
	DeletedAt               gorm.DeletedAt
}

type Autoclik_check_Story struct {
	Autoclik_check_StoryID string               `gorm:"type:uuid;primaryKey" json:"autoclik_check_story_id"`
	Autoclik_check_Name    string               `grom:"default:''" json:"autoclik_check_story_name"`
	Autoclik_check_details string               `grom:"default:''" json:"autoclik_check_story_details"`
	Item_No                *string              `json:"Item_No"`
	CreatedAt              *time.Time           `json:"created_at"`
	UpdatedAt              *time.Time           `json:"updated_at"`
	DeletedAt              gorm.DeletedAt       `gorm:"index"`
	UserID                 string               `json:"user_id"`
	User                   User                 `validate:"-"`
	Autoclik_Round_CountID string               `json:"autoclik_round_count_id"`
	Autoclik_Round_Count   Autoclik_Round_Count `validate:"-"`
}
