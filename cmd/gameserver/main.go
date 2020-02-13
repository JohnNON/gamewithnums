package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/JohnNON/gamewithnums/internal/app/gameserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/gameserver.toml", "путь до конфиг-файла")
}

func main() {
	flag.Parse()

	config := gameserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("PORT") != "" {
		config.BindAddr = ":" + os.Getenv("PORT")
	}

	if os.Getenv("LOG_LEVEL") != "" {
		config.LogLevel = os.Getenv("LOG_LEVEL")
	}

	if os.Getenv("DATABASE_URL") != "" {
		config.DatabaseURL = os.Getenv("DATABASE_URL")
	}

	if os.Getenv("DATABASE_DRIVER") != "" {
		config.DatabaseDriver = os.Getenv("DATABASE_DRIVER")
	}

	if os.Getenv("DATABASE_URL") != "" {
		config.DatabaseURL = os.Getenv("DATABASE_URL")
	}

	if os.Getenv("SESSION_KEY") != "" {
		config.SessionKey = os.Getenv("SESSION_KEY")
	}

	if os.Getenv("SESSION_MAXAGE") != "" {
		config.SessionMaxAge = 3600
		maxage, err := strconv.Atoi(os.Getenv("SESSION_MAXAGE"))
		if err == nil {
			config.SessionMaxAge = maxage
		}
	}

	if os.Getenv("CSRF_KEY") != "" {
		config.CsrfKey = os.Getenv("CSRF_KEY")
	}

	if err := gameserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
