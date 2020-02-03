package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер для подключения к PostgreSQL
)

// Store - структура описывающее хранилище
type Store struct {
	config         *Config
	db             *sqlx.DB
	userRepository *UserRepository
}

// New - функция создающее новое хранилище
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open - открывает соединение с хранилищем
func (s *Store) Open() error {
	db, err := sqlx.Connect(s.config.DatabaseDriver, s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close - закрывает соединение с хранилищем
func (s *Store) Close() {
	s.db.Close()
}

// User - метод для работы с репозиторием user
func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
