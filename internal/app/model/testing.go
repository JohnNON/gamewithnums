package model

import "testing"

// TestUser - вернет подготовленного User для тестов
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
