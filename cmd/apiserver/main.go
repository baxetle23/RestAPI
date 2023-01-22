package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"httpRestAPI/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
