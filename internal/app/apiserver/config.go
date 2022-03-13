package apiserver

import (
	"log"
	"os"

	"github.com/mserebryaakov/books-rest-api/internal/app/store"
)

// Config ...
type Config struct {
	BindAddr string `toml:"BindAddr"`
	Log      *log.Logger
	Store    *store.Config
}

// New config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Log:      log.New(os.Stdout, "books-api-debug", log.LstdFlags),
		Store:    store.NewConfig(),
	}
}
