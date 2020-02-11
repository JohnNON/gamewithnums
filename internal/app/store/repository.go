package store

import "github.com/JohnNON/gamewithnums/internal/app/model"

// UserRepository - интерфейс, описывающий хранилище пользователя
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

// RecordRepository - интерфейс, описывающий хранилище таблицы рекордов
type RecordRepository interface {
	Create(*model.Record) error
	FindByUserID(string, string) (*[]model.Record, error)
	GetAllRecords(string) (*[]model.Record, error)
}

// RoundRepository - интерфейс, описывающий хранилище таблицы рекордов
type RoundRepository interface {
	Create(*model.Round) error
	FindByUserID(string) (*[]model.Round, error)
	DeleteByUserID(string) error
}
