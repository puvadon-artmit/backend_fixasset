package model

import (
	"time"

	"gorm.io/gorm"
)

type Assets_Count_Store struct {
	Assets_Count_StoreID string  `gorm:"type:uuid;primaryKey" json:"assets_count_store_id"`
	Model_name           *string `grom:"default:''" json:"model_name"`
	Manufacturer         *string `grom:"default:''" json:"manufacturer"`
	Serial_Code          *string `grom:"default:''" json:"serial_code"`
	Type                 *string `grom:"default:''" json:"type"`
	Model                *string `grom:"default:''" json:"model"`
	Branch               *string `grom:"default:''" json:"branch"`
	Username             *string `grom:"default:''" json:"username"`
	Property_code        *string `grom:"default:''" json:"property_code"`
	Status               *string `grom:"default:''" json:"status"`
	Group_hardware       *string `grom:"default:''" json:"group_hardware"`
	Group                *string `grom:"default:''" json:"group"`
	User_hardware        *string `grom:"default:''" json:"user_hardware"`
	Phone_number         *string `grom:"default:''" json:"phone_number"`
	Posting_group        *string `grom:"default:''" json:"posting_group"`
	Latest_time          *string `grom:"default:''" json:"latest_time"`
	ResponsibleID        *string `grom:"default:''" json:"responsible_id"`
	// Responsible          Responsible    `validate:"-"`
	Comment1    *string        `grom:"default:''" json:"comment1"`
	Comment2    *string        `grom:"default:''" json:"comment2"`
	Comment3    *string        `grom:"default:''" json:"comment3"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      *string        `json:"user_id"`
	User        User           `validate:"-"`
	CategoryID  *string        `gorm:"default:NULL" json:"category_id"`
	Category    Category       `validate:"-"`
	ItemModelID *string        `grom:"default:''" json:"item_model_id"`
	// Item_model           Item_model     `validate:"-"`
	GroundID *string `grom:"default:''" json:"ground_id"`
	// Ground               Ground         `validate:"-"`
	Asset_countID *string     `json:"asset_count_id"`
	Asset_count   Asset_count `validate:"-"`
}
