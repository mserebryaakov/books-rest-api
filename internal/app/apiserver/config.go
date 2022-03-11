package apiserver

import (
	"log"
	"os"
)

// Config ...
type Config struct {
	BindAddr string
	Log      *log.Logger
}

// New config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Log:      log.New(os.Stdout, "books-api-debug", log.LstdFlags),
	}
}
