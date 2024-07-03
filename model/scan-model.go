package model

import (
	"time"

	"gorm.io/gorm"
)

type Scan_story struct {
	Scan_storyID  string         `gorm:"type:uuid;primaryKey" json:"scan_story_id"`
	Scan_name     *string        `json:"scan_name"`
	Scan_detail   *string        `json:"scan_detail"`
	Property_code string         `json:"property_code"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Round_CountID string         `json:"round_count_id"`
	Round_Count   Round_Count    `validate:"-"`
	UserID        string         `json:"user_id"`
	User          User           `validate:"-"`
}
