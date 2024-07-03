package model

import (
	"time"

	"gorm.io/gorm"
)

type Request_Update_Data struct {
	Request_Update_DataID string         `gorm:"type:uuid;primaryKey" json:"request_update_data_id"`
	Status                string         `json:"status"`
	Group_api             string         `json:"group_api"`
	CreatedAt             *time.Time     `json:"created_at"`
	UpdatedAt             *time.Time     `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
