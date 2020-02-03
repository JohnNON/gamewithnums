package main

import (
	"flag"
	"log"

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

	server := gameserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
