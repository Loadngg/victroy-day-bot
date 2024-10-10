package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func MustLoad() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TELEGRAM_APITOKEN")
	cfg := Config{Token: token}

	return &cfg
}
