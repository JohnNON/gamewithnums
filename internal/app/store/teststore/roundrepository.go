package teststore

import (
	"strconv"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// RoundRepository - структура описывает хранилище для тестирования
type RoundRepository struct {
	store  *Store
	rounds map[int]*model.Round
}

// Create - метода создает запись пользователя в тестовом хранилище
func (r *RoundRepository) Create(rn *model.Round) error {
	if err := rn.Validate(); err != nil {
		return err
	}

	rn.ID = len(r.rounds) + 1
	r.rounds[rn.ID] = rn
	rn.ID = len(r.rounds)

	return nil
}

// FindByUserID - ищет user по значению поля userID в тестовом хранилище
func (r *RoundRepository) FindByUserID(userID string) (*[]model.Round, error) {
	rounds := &[]model.Round{}
	for i, rn := range r.rounds {
		if i == 10 {
			break
		}
		if strconv.Itoa(rn.UserID) == userID {
			*rounds = append(*rounds, *rn)
		}
	}

	if len(*rounds) > 0 {
		return rounds, nil
	}
	return nil, store.ErrRecordNotFound
}

// DeleteByUserID - ищет user по значению поля userID в тестовом хранилище
func (r *RoundRepository) DeleteByUserID(userID string) error {
	r.rounds = make(map[int]*model.Round)
	return nil
}
