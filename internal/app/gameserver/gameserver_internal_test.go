package gameserver

import "testing"

import "net/http/httptest"

import "net/http"

import "github.com/stretchr/testify/assert"

func TestGameServer_handleHello(t *testing.T) {
	server := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	server.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello!")
}
