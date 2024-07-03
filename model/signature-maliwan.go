package model

import (
	"time"

	"gorm.io/gorm"
)

type Signature_Maliwan struct {
	Signature_MaliwanID string         `gorm:"type:uuid;primaryKey" json:"maliwan_signature_id"`
	Maliwan_Signature   *string        `json:"maliwan_signature"`
	UserID              string         `json:"user_id"`
	User                User           `validate:"-"`
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	Maliwan_countID     string         `json:"maliwan_count_id"`
	Maliwan_count       Maliwan_count  `validate:"-"`
}
