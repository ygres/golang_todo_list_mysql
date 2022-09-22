package main

import (
	"app/internal/app/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPah string
)

func init() {
	flag.StringVar(&configPah, "config", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := &apiserver.Config{}
	_, err := toml.DecodeFile(configPah, config)

	if err != nil {
		log.Fatal(err)
	}

	server := &apiserver.Server{}
	server.Initialize(config.DatabaseURL)
	server.Start(config.BindAddr)
}
