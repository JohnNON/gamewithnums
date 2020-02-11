package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Record - описание модели таблици рекордов
type Record struct {
	ID         int
	UserID     int
	Difficulty int
	RoundCount int
	GameTime   int
	User
}

// Validate - метод для валидации вводимых данных
func (r *Record) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.UserID, validation.Required, validation.Min(1)),
		validation.Field(&r.Difficulty, validation.Min(1)),
		validation.Field(&r.RoundCount, validation.Min(1)),
		validation.Field(&r.GameTime, validation.Min(1)),
	)
}
