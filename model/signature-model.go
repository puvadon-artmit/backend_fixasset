package model

import (
	"time"

	"gorm.io/gorm"
)

type Signature struct {
	SignatureID   string         `gorm:"type:uuid;primaryKey" json:"signature_id"`
	Signature     *string        `json:"signature"`
	UserID        string         `json:"user_id"`
	User          User           `validate:"-"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Asset_countID string         `json:"asset_count_id"`
	Asset_count   Asset_count    `validate:"-"`
}
