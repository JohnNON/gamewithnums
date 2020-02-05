package gameserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JohnNON/gamewithnums/internal/app/model"
	"github.com/JohnNON/gamewithnums/internal/app/store/teststore"
	"github.com/gorilla/securecookie"
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
			req := httptest.NewRequest(http.MethodPost, "/users", b)
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
			req := httptest.NewRequest(http.MethodPost, "/sessions", b)
			server.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestGameserver_authenticateUser(t *testing.T) {
	store := teststore.New()
	u := model.TestUser(t)
	if err := store.User().Create(u); err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name         string
		coockieVal   map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			coockieVal: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			coockieVal:   nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			cookieStr, err := sc.Encode(sessionName, tc.coockieVal)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
