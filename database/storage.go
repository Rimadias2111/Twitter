package database

import (
	"gorm.io/gorm"
	"project/database/storage"
	"project/models"
)

type IStore interface {
	User() *User
}

type Store struct {
	db   *gorm.DB
	user User
}

type User interface {
	CreateUser(user *models.User) (string, error)
}

func New(db *gorm.DB) {
	store := &Store{}

	store.user = storage.NewUserRepo(db)

}

func (s *Store) User() *User { return &s.user }

func test(s *Store) {

}
