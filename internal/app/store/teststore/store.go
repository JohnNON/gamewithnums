package teststore

import (
	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// Store - структура описывающее хранилище
type Store struct {
	userRepository   *UserRepository
	recordRepository *RecordRepository
	roundRepository  *RoundRepository
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
			users: make(map[int]*model.User),
		}
	}

	return s.userRepository
}

// Record - метод для работы с репозиторием record
func (s *Store) Record() store.RecordRepository {
	if s.recordRepository == nil {
		s.recordRepository = &RecordRepository{
			store:   s,
			records: make(map[int]*model.Record),
		}
	}

	return s.recordRepository
}

// Round - метод для работы с репозиторием round
func (s *Store) Round() store.RoundRepository {
	if s.roundRepository == nil {
		s.roundRepository = &RoundRepository{
			store:  s,
			rounds: make(map[int]*model.Round),
		}
	}

	return s.roundRepository
}
