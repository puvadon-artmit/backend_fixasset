package model

import (
	"time"

	"gorm.io/gorm"
)

type Round_Count struct {
	Round_CountID string         `gorm:"type:uuid;primaryKey" json:"round_count_id"`
	Round         *string        `json:"round"`
	Status        *string        `json:"status"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Asset_countID string         `json:"asset_count_id"`
	Asset_count   Asset_count    `validate:"-"`
}

type Round_Count_Story struct {
	Round_Count_StoryID string         `gorm:"type:uuid;primaryKey" json:"round_count_story_id"`
	Round_Name_Strory   *string        `json:"round_name_strory"`
	Round_Name          *string        `json:"round_name"`
	Round_Strory_Detail *string        `json:"round_strory_detail"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	Round_CountID       string         `json:"round_count_id"`
	Round_Count         Round_Count    `validate:"-"`
	UserID              string         `json:"user_id"`
	User                User           `validate:"-"`
}
