package model

import (
	"time"

	"gorm.io/gorm"
)

type Responsible struct {
	ResponsibleID   string         `gorm:"type:uuid;primaryKey" json:"responsible_id"`
	ResponsibleName *string        `grom:"default:''" json:"responsible_name"`
	EmployeeCode    *string        `grom:"default:''" json:"employee_code"`
	Comment         *string        `grom:"default:''" json:"comment"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          string         `json:"user_id"`
	GroupID         string         `json:"group_id"`
	Group           Group          `validate:"-"`
}

type Responsible_Story struct {
	ResponsibleStoryID       string         `gorm:"type:uuid;primaryKey" json:"responsible_story_id"`
	Responsible_StoryName    *string        `grom:"default:''" json:"responsible_story_name"`
	Responsible_StoryDetails *string        `grom:"default:''" json:"responsible_story_details"`
	CreatedAt                *time.Time     `json:"created_at"`
	UpdatedAt                *time.Time     `json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index"`
	UserID                   string         `json:"user_id"`
	User                     User           `validate:"-"`
	ResponsibleID            string         `json:"responsible_id"`
	Responsible              Responsible    `validate:"-"`
}
