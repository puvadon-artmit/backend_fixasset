package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID       string         `gorm:"type:uuid;primaryKey" json:"user_id"`
	Firstname    string         `json:"firstname"`
	Lastname     string         `json:"lastname" `
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	Status       string         `json:"status"`
	EmployeeCode string         `json:"employee_code"`
	RoleActive   *string        `json:"role_active"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Roles        []Role         `gorm:"many2many:user_roles;"`
}

// เอาไว้เก็บ UserID และ RoleID เพื่อ ให้ 1 user มีได้หลาย Role
type UserRoles struct {
	UserID string `gorm:"type:uuid"`
	RoleID string `gorm:"type:uuid"`
}

// เอาไว้เก็บ Role ตำแหน่งต่างๆ
type Role struct {
	RoleID          string            `gorm:"type:uuid;primaryKey" json:"role_id"`
	RoleName        string            `json:"role_name" validate:"required"`
	RoleDisplayName string            `json:"role_display_name" validate:"required"`
	RoleDescription string            `json:"role_description" validate:"required"`
	RoleCode        string            `json:"role_code"`
	Status          bool              `json:"status" gorm:"default:true"`
	CreatedAt       *time.Time        `json:"created_at"`
	UpdatedAt       *time.Time        `json:"updated_at"`
	Users           []User            `gorm:"many2many:user_roles;"`
	PermissionGroup []PermissionGroup `gorm:"foreignKey:RoleID"`
}

// เอาไว้เก็บ การทำงานหลักๆ ว่า สามารถทำอะไรได้
type Permission struct {
	PermissionID          string                `gorm:"type:uuid;primaryKey" json:"permission_id"`
	PermissionName        string                `json:"permission_name" validate:"required"`
	PermissionDisplayName string                `json:"permission_display_name" validate:"required"`
	PermissionDescription string                `json:"permission_description" validate:"required"`
	Status                bool                  `json:"status" gorm:"default:true"`
	CreatedAt             *time.Time            `json:"created_at"`
	UpdatedAt             *time.Time            `json:"updated_at"`
	PermissionComponent   []PermissionComponent `gorm:"foreignKey:PermissionID"`
	PermissionGroup       []PermissionGroup     `gorm:"foreignKey:PermissionID"`
}

// เอาไว้เก็บ การทำงานแต่ล่ะ Component โดยละเอียดยิบ
type PermissionComponent struct {
	PermissionComponentID          string            `gorm:"type:uuid;primaryKey" json:"permission_component_id"`
	PermissionComponentName        string            `json:"permission_component_name"`
	PermissionComponentDisplayName string            `json:"permission_component_display_name"`
	PermissionComponentDescription string            `json:"permission_component_description"`
	Status                         bool              `json:"status" gorm:"default:true"`
	CreatedAt                      *time.Time        `json:"created_at"`
	UpdatedAt                      *time.Time        `json:"updated_at"`
	PermissionID                   string            `json:"permission_id"`
	Permission                     Permission        `validate:"-"`
	PermissionGroup                []PermissionGroup `gorm:"foreignKey:PermissionComponentID"`
}

// เอาไว้เก็บ การทำงานโดยรวมทุกอย่าง เพื่อส่งไปให้เลือก Role
type PermissionGroup struct {
	PermissionGroupID     string              `gorm:"type:uuid;primaryKey" json:"permission_group_id"`
	Activate              bool                `json:"avtivates"`
	CreatedAt             *time.Time          `json:"created_at"`
	UpdatedAt             *time.Time          `json:"updated_at"`
	RoleID                string              `json:"role_id"`
	PermissionID          string              `json:"permission_id"`
	PermissionComponentID string              `json:"permission_component_id"`
	Role                  Role                `validate:"-"`
	Permission            Permission          `validate:"-"`
	PermissionComponent   PermissionComponent `validate:"-"`
}
