package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/database"
	"project/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) database.User {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *models.User) (string, error) {
	err := r.db.Create(user).Error
	return uuid.New().String(), err
}
