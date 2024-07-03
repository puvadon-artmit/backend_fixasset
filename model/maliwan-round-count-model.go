package model

import (
	"time"

	"gorm.io/gorm"
)

type Maliwan_Round_Count struct {
	Maliwan_Round_CountID string         `gorm:"type:uuid;primaryKey" json:"maliwan_round_count_id"`
	Round                 *string        `json:"round"`
	Status                *string        `json:"status"`
	CreatedAt             *time.Time     `json:"created_at"`
	UpdatedAt             *time.Time     `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	Maliwan_countID       string         `json:"maliwan_count_id"`
	Maliwan_count         Maliwan_count  `validate:"-"`
}

type Maliwan_Round_Count_Story struct {
	Maliwan_Round_Count_StoryID      string              `gorm:"type:uuid;primaryKey" json:"maliwan_round_count_story_id"`
	Maliwan_Round_Count_Story_Name   string              `json:"maliwan_round_count_story_name"`
	Maliwan_Round_Count_Story_Detail string              `json:"maliwan_round_count_story_detail"`
	CreatedAt                        *time.Time          `json:"created_at"`
	UpdatedAt                        *time.Time          `json:"updated_at"`
	DeletedAt                        gorm.DeletedAt      `gorm:"index"`
	Maliwan_Round_CountID            string              `json:"maliwan_round_count_id"`
	Maliwan_Round_Count              Maliwan_Round_Count `validate:"-"`
	UserID                           string              `json:"user_id"`
	User                             User                `validate:"-"`
}
