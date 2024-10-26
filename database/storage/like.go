package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/database"
	"project/models"
)

type LikeRepo struct {
	db *gorm.DB
}

func NewLikeRepo(db *gorm.DB) database.Like {
	return &LikeRepo{
		db: db,
	}
}

func (r *LikeRepo) Create(like *models.Like) error {
	return r.db.Create(like).Error
}

func (r *LikeRepo) Delete(userID, tweetID uuid.UUID) error {
	return r.db.Where("user_id = ? AND tweet_id = ?", userID, tweetID).Delete(&models.Like{}).Error
}
