package teststore

import (
	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// UserRepository - структура описывает хранилище для тестирования
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create - метода создает запись пользователя в тестовом хранилище
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u
	u.ID = len(r.users)

	return nil
}

// Find - ищет user по значению поля email в тестовом хранилище
func (r *UserRepository) Find(id int) (*model.User, error) {
	if u, ok := r.users[id]; ok {
		u.Password = ""
		return u, nil
	}

	return nil, store.ErrRecordNotFound
}

// FindByEmail - ищет user по значению поля email в тестовом хранилище
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			u.Password = ""
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
