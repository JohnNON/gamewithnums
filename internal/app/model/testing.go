package model

import (
	"testing"
	"time"
)

// TestUser - вернет подготовленного User для тестов
func TestUser(t *testing.T) *User {
	return &User{
		Nickname: "user",
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

// TestRound - вернет подготовленного Round для тестов
func TestRound(t *testing.T) *Round {
	return &Round{
		UserID:     1,
		Difficulty: 1,
		GameNumber: "0987",
		GameTime:   time.Now().UTC().Format("2006-01-02 15:04:05"),
		Inpt:       "0872",
		Outpt:      "BKKX",
	}
}
