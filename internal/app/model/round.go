package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Round - описание модели временной таблицы результатов раундов
type Round struct {
	ID         int    `json:"-"`
	UserID     int    `json:"-"`
	Difficulty int    `json:"diff"`
	GameNumber string `json:"-"`
	GameTime   string `json:"-"`
	Inpt       string `json:"in"`
	Outpt      string `json:"out"`
}

// Validate - метод для валидации вводимых данных
func (r *Round) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.UserID, validation.Required, validation.Min(1)),
		validation.Field(&r.Difficulty, validation.Min(1)),
		validation.Field(&r.GameNumber, is.Digit),
		validation.Field(&r.GameTime, is.ASCII),
		validation.Field(&r.Inpt, is.Digit),
		validation.Field(&r.Outpt, is.Alpha),
	)
}
