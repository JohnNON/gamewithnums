package gameserver

// Config - содержит конфигурацию для запуска сервера
type Config struct {
	BindAddr       string `toml:"bind_addr"`
	LogLevel       string `toml:"log_level"`
	DatabaseURL    string `toml:"database_url"`
	DatabaseDriver string `toml:"database_driver"`
	SessionKey     string `toml:"session_key"`
}

// NewConfig - инициализация конфига по умолчанию
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
