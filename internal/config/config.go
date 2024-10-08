package config

import "os"

type Config struct {
	Token string
}

func MustLoad() *Config {
	cfg := Config{Token: os.Getenv("TELEGRAM_APITOKEN")}

	return &cfg
}
