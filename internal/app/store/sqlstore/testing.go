package sqlstore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
)

// TestDB - вспомогательная функция для тестирования работы с хранилищем
func TestDB(t *testing.T, databaseDriver, databaseURL string) (*sqlx.DB, func(...string)) {
	t.Helper()

	db, err := sqlx.Connect(databaseDriver, databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		db.Close()
	}
}
