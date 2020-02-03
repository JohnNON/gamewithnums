package model

// User - описание модели пользователя
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}
