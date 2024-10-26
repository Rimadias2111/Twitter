package database

import (
	"gorm.io/gorm"
	"project/database/storage"
)

type IStore interface {
	User() User
	Tweet() Tweet
	Like() Like
	Follow() Follow
}

type Store struct {
	db     *gorm.DB
	user   User
	tweet  Tweet
	like   Like
	follow Follow
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

func (s *Store) User() User { return s.user }

func (s *Store) Tweet() Tweet { return s.tweet }

func (s *Store) Like() Like { return s.like }

func (s *Store) Follow() Follow { return s.follow }
