package teststore

import (
	"strconv"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store"
)

// RecordRepository - структура описывает хранилище для тестирования
type RecordRepository struct {
	store   *Store
	records map[int]*model.Record
}

// Create - метода создает запись пользователя в тестовом хранилище
func (r *RecordRepository) Create(rc *model.Record) error {
	if err := rc.Validate(); err != nil {
		return err
	}

	rc.ID = len(r.records) + 1
	r.records[rc.ID] = rc
	rc.ID = len(r.records)

	return nil
}

// FindByUserID - ищет user по значению поля email в тестовом хранилище
func (r *RecordRepository) FindByUserID(userID string, diff string) (*[]model.Record, error) {
	records := &[]model.Record{}
	for i, rc := range r.records {
		if i == 10 {
			break
		}
		if strconv.Itoa(rc.UserID) == userID && strconv.Itoa(rc.Difficulty) == diff {
			*records = append(*records, *rc)
		}
	}

	if len(*records) > 0 {
		return records, nil
	}
	return nil, store.ErrRecordNotFound
}

// GetAllRecords - ищет user по значению поля email в тестовом хранилище
func (r *RecordRepository) GetAllRecords(diff string) (*[]model.Record, error) {
	records := &[]model.Record{}
	for i, rc := range r.records {
		if i == 10 {
			break
		}
		*records = append(*records, *rc)
	}

	if len(*records) > 0 {
		return records, nil
	}
	return nil, store.ErrRecordNotFound
}
