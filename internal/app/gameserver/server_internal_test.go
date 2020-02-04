package gameserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestGameServer_handleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	server := newServer(teststore.New())
	server.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Equal(t, rec.Body.String(), "Hello Users!")
}
