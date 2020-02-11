package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL, databaseDriver string
)

// TestMain - запустится перед тестами в store_test один раз, используется для инициализации конфигурации
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	databaseDriver = os.Getenv("DATABASE_DRIVER")
	if databaseURL == "" || databaseDriver == "" {
		databaseURL = "user=postgres password=C0nf1cer dbname=gamewithnums sslmode=disable"
		databaseDriver = "postgres"
	}

	os.Exit(m.Run())
}
