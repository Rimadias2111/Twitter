package models

import (
	"github.com/google/uuid"
	"time"
)

type Follow struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	FollowerID uuid.UUID `gorm:"type:uuid;not null"`
	FollowedID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
