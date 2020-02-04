package teststore

import (
	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// Store - структура описывающее хранилище
type Store struct {
	userRepository *UserRepository
}

// New - функция создающее новое хранилище
func New() *Store {
	return &Store{}
}

// User - метод для работы с репозиторием user
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}

	return s.userRepository
}
