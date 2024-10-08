package main

import (
	"log"

	"victory-day-bot/internal/config"
	"victory-day-bot/internal/pkg/app"
)

func main() {
	cfg := config.MustLoad()
	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
