package model

import "testing"

// TestUser - вернет подготовленного User для тестов
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

// TestRecord - вернет подготовленного Record для тестов
func TestRecord(t *testing.T) *Record {
	return &Record{
		UserID:     1,
		Difficulty: 1,
		RoundCount: 10,
		GameTime:   120,
	}
}
