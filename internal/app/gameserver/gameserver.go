package gameserver

import (
	"net/http"

	"github.com/JohnNON/gamewithnums/internal/app/store/sqlstore"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

// Start - выполняет запуск сервера
func Start(config *Config) error {
	db, err := newDB(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	sessionStore.MaxAge(config.SessionMaxAge)

	srv := newServer(store, sessionStore)

	CSRF := csrf.Protect([]byte(config.CsrfKey), csrf.Secure(false), csrf.FieldName("Csrf"))

	return http.ListenAndServe(config.BindAddr, CSRF(srv))
}

func newDB(databaseDriver, databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open(databaseDriver, databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
