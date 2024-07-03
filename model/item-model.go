package model

import (
	"time"

	"gorm.io/gorm"
)

type Item_model struct {
	ItemModelID   string         `gorm:"type:uuid;primaryKey" json:"item_model_id"`
	ItemModelName *string        `json:"item_model_name" validate:"required"`
	Comment       *string        `json:"comment"`
	ProductNumber *string        `json:"product_number"`
	Weight        *string        `json:"weight"`
	RequiredUnits *int           `json:"required_units"`
	Frontpicture  *string        `json:"frontpicture"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string         `json:"user_id"`
	TypeID        string         `json:"type_id"`
	Type          Type_things    `validate:"-"`
}

type Story struct {
	StoryID     string         `gorm:"type:uuid;primaryKey" json:"story_id"`
	StoryName   *string        `json:"story_name" validate:"required"`
	Details     *string        `json:"details"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      string         `json:"user_id"`
	User        User           `validate:"-"`
	ItemModelID string         `json:"item_model_id"`
	Item_model  Item_model     `validate:"-"`
}
