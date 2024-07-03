package model

import (
	"time"

	"gorm.io/gorm"
)

type Signature_Autoclik struct {
	Signature_AutoclikID string         `gorm:"type:uuid;primaryKey" json:"autoclik_signature_id"`
	Autoclik_Signature   *string        `json:"autoclik_signature"`
	UserID               string         `json:"user_id"`
	User                 User           `validate:"-"`
	CreatedAt            *time.Time     `json:"created_at"`
	UpdatedAt            *time.Time     `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	Autoclik_countID     string         `json:"autoclik_count_id"`
	Autoclik_count       Autoclik_count `validate:"-"`
}
