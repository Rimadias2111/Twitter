package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/models"
)

type FollowRepo struct {
	db *gorm.DB
}

func NewFollowRepo(db *gorm.DB) Follow {
	return &FollowRepo{db: db}
}

func (r *FollowRepo) Create(follow *models.Follow) error {
	return r.db.Create(follow).Error
}

func (r *FollowRepo) Delete(followerID, followedID uuid.UUID) error {
	return r.db.Where("follower_id = ? AND followed_id = ?", followerID, followedID).Delete(&models.Follow{}).Error
}

func (r *FollowRepo) IsFollowing(followerID, followedID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Follow{}).
		Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Count(&count).Error
	return count > 0, err
}
