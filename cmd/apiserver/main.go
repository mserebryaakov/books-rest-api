package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/mserebryaakov/books-rest-api/internal/app/apiserver"
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

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		config.Log.Fatal(err)
	}

	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		config.Log.Fatal(err)
	}
}
