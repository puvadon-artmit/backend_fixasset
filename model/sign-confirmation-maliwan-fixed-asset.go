package model

import (
	"time"

	"gorm.io/gorm"
)

type Signature_Maliwan_Fixed_Asset struct {
	Signature_Maliwan_Fixed_AssetID string                    `gorm:"type:uuid;primaryKey" json:"maliwan_signature_id"`
	Maliwan_Fixed_Asset_Signature   *string                   `json:"maliwan_fixed_asset_signature"`
	UserID                          string                    `json:"user_id"`
	User                            User                      `validate:"-"`
	CreatedAt                       *time.Time                `json:"created_at"`
	UpdatedAt                       *time.Time                `json:"updated_at"`
	DeletedAt                       gorm.DeletedAt            `gorm:"index"`
	Maliwan_Fixed_Asset_CountID     string                    `json:"maliwan_fixed_asset_count_id"`
	Maliwan_Fixed_Asset_Count       Maliwan_Fixed_Asset_Count `validate:"-"`
}
