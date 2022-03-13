package store

type Config struct {
	DataBaseURL string `toml:"DataBaseURL"`
}

func NewConfig() *Config {
	return &Config{}
}
