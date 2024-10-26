package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Tweet struct {
	Id        uuid.UUID  `gorm:"primary_key; type:uuid"`
	UserID    uuid.UUID  `gorm:"type:uuid; not null; foreign_key; references: user_id; constraint: OnUpdate:CASCADE, OnDelete: SET NULL"`
	Content   string     `gorm:"type:text; not null"`
	ImagePath *string    `gorm:"size:255"`
	VideoPath *string    `gorm:"size:255"`
	RetweetID *uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetAllTweetsRequest struct {
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
	UserID string `json:"user_id"`
	Search string `json:"search"`
}

type GetAllTweetsResponse struct {
	Tweets []Tweet `json:"tweets"`
	Count  int64   `json:"count"`
}
