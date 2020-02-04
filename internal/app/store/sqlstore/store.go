package sqlstore

import (
	"github.com/JohnNON/gamewithnums/internal/app/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер для подключения к PostgreSQL
)

// Store - структура описывающее хранилище
type Store struct {
	db             *sqlx.DB
	userRepository *UserRepository
}

// New - функция создающее новое хранилище
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// User - метод для работы с репозиторием user
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
