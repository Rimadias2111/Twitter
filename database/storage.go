package database

import (
	"gorm.io/gorm"
	"project/database/storage"
)

type IStore interface {
	User() storage.User
	Tweet() storage.Tweet
	Like() storage.Like
	Follow() storage.Follow
}

type Store struct {
	db     *gorm.DB
	user   storage.User
	tweet  storage.Tweet
	like   storage.Like
	follow storage.Follow
}

func New(db *gorm.DB) *Store {
	return &Store{
		db:     db,
		user:   storage.NewUserRepo(db),
		tweet:  storage.NewTweetRepo(db),
		like:   storage.NewLikeRepo(db),
		follow: storage.NewFollowRepo(db),
	}
}

func (s *Store) User() storage.User { return s.user }

func (s *Store) Tweet() storage.Tweet { return s.tweet }

func (s *Store) Like() storage.Like { return s.like }

func (s *Store) Follow() storage.Follow { return s.follow }
