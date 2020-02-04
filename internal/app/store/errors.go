package store

import "errors"

var (
	// ErrRecordNotFound - ошибка, возвращаемая при отсутствии записи в хранилище
	ErrRecordNotFound = errors.New("record not found")
)
