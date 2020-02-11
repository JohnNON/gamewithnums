package main

import (
	"flag"
	"log"
	"os"

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

	if err := gameserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
