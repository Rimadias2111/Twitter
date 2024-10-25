package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"primary_key; type:uuid;"`
	Name     string    `gorm:"size:255; not null"`
	Username string    `gorm:"size:255; unique; not null"`
}
