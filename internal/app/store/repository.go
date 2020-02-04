package store

import "github.com/JohnNON/gamewithnums/internal/app/model"

// UserRepository - интерфейс, описывающий хранилище пользователя
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
