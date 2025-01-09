package usersmasterdata

import (
	"encoding/json"
	"time"
)

type User struct {
	ID           uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	SchoolID     uint64          `gorm:"not null" json:"school_id"`
	Email        string          `gorm:"size:255;unique" json:"email"`
	Phone        string          `gorm:"size:20;unique" json:"phone"`
	SocialLogins json.RawMessage `gorm:"type:jsonb" json:"social_logins"`
	DeletedBy    *uint64         `json:"deleted_by"`
	DeletedAt    *time.Time      `json:"deleted_at"`
	CreatedAt    time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

type Role struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SchoolID  uint64    `gorm:"not null" json:"school_id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Permission struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type RolePermission struct {
	RoleID       uint64    `gorm:"primaryKey" json:"role_id"`
	PermissionID uint64    `gorm:"primaryKey" json:"permission_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type UserRole struct {
	UserID    uint64    `gorm:"primaryKey" json:"user_id"`
	RoleID    uint64    `gorm:"primaryKey" json:"role_id"`
	SchoolID  uint64    `gorm:"primaryKey" json:"school_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
