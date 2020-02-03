package store

// Config - содержит конфигурацию для работы с хранилищем
type Config struct {
	DatabaseURL    string `toml:"database_url"`
	DatabaseDriver string `toml:"database_driver"`
}

// NewConfig - возвращает базовую конфигурацию хранилища
func NewConfig() *Config {
	return &Config{
		DatabaseURL:    "user=postgres dbname=gamewithnums sslmode=disable",
		DatabaseDriver: "postgres",
	}
}
