package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           uuid.UUID `gorm:"primary_key; type:uuid;"`
	Name         string    `gorm:"size:255; not null"`
	Bio          *string   `gorm:"size:255;"`
	Username     string    `gorm:"size:255; unique; not null; uniqueIndex:idx_username_deleted_at"`
	Password     string    `gorm:"size:255; not null"`
	ProfileImage *string   `gorm:"size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index; uniqueIndex:idx_username_deleted_at"`
}

type GetAllUsersRequest struct {
	Page      uint64    `json:"page"`
	Limit     uint64    `json:"limit"`
	Search    string    `json:"search"`
	Followers uuid.UUID `json:"id_followers"`
	Following uuid.UUID `json:"id_following"`
}

type GetAllUsersResponse struct {
	Users []User `json:"users"`
	Count int64  `json:"count"`
}

type CreateUser struct {
	Name         string  `json:"name"`
	Bio          *string `json:"bio"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	ProfileImage *string `json:"profileImage"`
}

type UpdateUser struct {
	Name         string  `json:"name"`
	Bio          *string `json:"bio"`
	Username     string  `json:"username"`
	ProfileImage *string `json:"profileImage"`
}
