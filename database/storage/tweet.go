package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/database"
	"project/models"
)

type TweetRepo struct {
	db *gorm.DB
}

func NewTweetRepo(db *gorm.DB) database.Tweet {
	return &TweetRepo{
		db: db,
	}
}

func (r *TweetRepo) Create(tweet *models.Tweet) (string, error) {
	id := uuid.New()
	tweet.Id = id
	if err := r.db.Create(tweet).Error; err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *TweetRepo) Update(tweet *models.Tweet) error {
	if err := r.db.Save(tweet).Error; err != nil {
		return err
	}
	return nil
}

func (r *TweetRepo) Delete(req models.RequestId) error {
	if err := r.db.Where("id = ?", req.Id).Delete(&models.Tweet{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *TweetRepo) Get(req models.RequestId) (*models.Tweet, error) {
	var tweet models.Tweet
	if err := r.db.Where("id = ?", req.Id).First(&tweet).Error; err != nil {
		return nil, err
	}
	return &tweet, nil
}

func (r *TweetRepo) GetAll(req models.GetAllTweetsRequest) (*models.GetAllTweetsResponse, error) {
	var (
		resp   models.GetAllTweetsResponse
		query  = r.db.Model(&models.Tweet{})
		offset = (req.Page - 1) * req.Limit
	)

	if req.Search != "" {
		query = query.Where("content ILIKE ?", "%"+req.Search+"%")
	}

	if req.UserID != "" {
		query = query.Where("user_id = ?", req.UserID)
	}

	if err := query.Offset(int(offset)).Limit(int(req.Limit)).Find(&resp.Tweets).Error; err != nil {
		return nil, err
	}

	if err := query.Count(&resp.Count).Error; err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *TweetRepo) GetTweetsForUser(userID uuid.UUID, req models.GetAllTweetsRequest) (*models.GetAllTweetsResponse, error) {
	var (
		resp   models.GetAllTweetsResponse
		offset = int((req.Page - 1) * req.Limit)
	)

	subQuery := r.db.Model(&models.Follow{}).Select("followed_id").Where("follower_id = ?", userID)

	query := r.db.Model(&models.Tweet{}).
		Where("user_id IN (?)", subQuery).
		Or("user_id = ?", userID).
		Offset(offset).
		Limit(int(req.Limit)).
		Order("created_at DESC")

	if err := query.Find(&resp.Tweets).Error; err != nil {
		return nil, err
	}

	if err := query.Count(&resp.Count).Error; err != nil {
		return nil, err
	}

	return &resp, nil
}
