package gameserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestGameServer_handleUsersCreate(t *testing.T) {
	server := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret key")))
	u := model.TestUser(t)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "user---example.org",
				"password": u.Password,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "empty email",
			payload: map[string]string{
				"email":    "",
				"password": u.Password,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "pass",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "empty password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "invalid payload",
			payload:      "payload",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			server.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestGameserver_handleSessionsCreate(t *testing.T) {
	server := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret key")))
	u := model.TestUser(t)
	if err := server.store.User().Create(u); err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "us@ex.org",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "pass",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "invalid payload",
			payload:      "payload",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			server.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
