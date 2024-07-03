package model

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	LocationID    string `gorm:"type:uuid;primaryKey" json:"location_id"`
	Coordinates_x string `json:"coordinates_x"`
	Coordinates_y string `json:"coordinates_y"`
	AssetsID      string `json:"assets_id"`
	Assets        Assets `gorm:"foreignKey:AssetsID" validate:"-"`
	GroundID      string `json:"ground_id"`
	Ground        Ground `gorm:"foreignKey:GroundID" validate:"-"`
}
type Location_story struct {
	Location_storyID   string         `gorm:"type:uuid;primaryKey" json:"location_story_id"`
	Location_StoryName *string        `json:"location_story_name"`
	Location_Details   *string        `json:"location_details"`
	CreatedAt          *time.Time     `json:"created_at"`
	UpdatedAt          *time.Time     `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	LocationID         string         `json:"location_id"`
	Location           Location       `validate:"-"`
	UserID             string         `json:"user_id"`
	User               User           `validate:"-"`
}
