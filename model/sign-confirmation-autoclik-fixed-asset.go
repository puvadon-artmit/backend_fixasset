package model

import (
	"time"

	"gorm.io/gorm"
)

type Signature_Autoclik_Fixed_Asset struct {
	Signature_Autoclik_Fixed_AssetID string                     `gorm:"type:uuid;primaryKey" json:"autoclik_signature_id"`
	Autoclik_Fixed_Asset_Signature   *string                    `json:"autoclik_fixed_asset_signature"`
	UserID                           string                     `json:"user_id"`
	User                             User                       `validate:"-"`
	CreatedAt                        *time.Time                 `json:"created_at"`
	UpdatedAt                        *time.Time                 `json:"updated_at"`
	DeletedAt                        gorm.DeletedAt             `gorm:"index"`
	Autoclik_Fixed_Asset_CountID     string                     `json:"autoclik_fixed_asset_count_id"`
	Autoclik_Fixed_Asset_Count       Autoclik_Fixed_Asset_Count `validate:"-"`
}
