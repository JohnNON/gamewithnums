package teststore

import (
	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// UserRepository - структура описывает хранилище для тестирования
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create - метода создает запись пользователя в тестовом хранилище
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail - ищет user по значению поля email в тестовом хранилище
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	if u, ok := r.users[email]; ok {
		u.Password = ""
		return u, nil
	}

	return nil, store.ErrRecordNotFound
}
