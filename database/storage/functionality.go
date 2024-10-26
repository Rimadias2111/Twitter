package storage

import (
	"github.com/google/uuid"
	"project/models"
)

type User interface {
	Create(user *models.User) (string, error)
	Update(user *models.User) error
	Delete(req models.RequestId) error
	Get(req models.RequestId) (*models.User, error)
	GetAll(req models.GetAllUsersRequest) (*models.GetAllUsersResponse, error)
	GetByUsername(username string) (*models.User, error)
}

type Tweet interface {
	Create(tweet *models.Tweet) (string, error)
	Update(tweet *models.Tweet) error
	Delete(req models.RequestId) error
	Get(req models.RequestId) (*models.Tweet, error)
	GetAll(req models.GetAllTweetsRequest) (*models.GetAllTweetsResponse, error)
	GetTweetsForUser(userID uuid.UUID, req models.GetAllTweetsRequest) (*models.GetAllTweetsResponse, error)
}

type Like interface {
	Create(like *models.Like) error
	Delete(userID, tweetID uuid.UUID) error
}

type Follow interface {
	Create(follow *models.Follow) error
	Delete(followerID, followedID uuid.UUID) error
	IsFollowing(followerID, followedID uuid.UUID) (bool, error)
}
