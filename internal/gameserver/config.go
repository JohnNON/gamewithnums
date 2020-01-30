package gameserver

// Config - содержит конфигурацию для запуска сервера
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig - инициализация конфига по умолчанию
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
