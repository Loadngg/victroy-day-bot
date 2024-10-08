package app

import (
	"fmt"

	"victory-day-bot/internal/app/command"
	"victory-day-bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	bot      *tgbotapi.BotAPI
	commands *command.Command
}

func New(cfg *config.Config) (*App, error) {
	app := &App{}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	app.bot = bot
	if err != nil {
		return nil, err
	}

	app.bot.Debug = true

	app.commands = command.New()

	return app, nil
}

func (app *App) Run() error {
	fmt.Println("bot started")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := app.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg = app.commands.Start(&msg)
		default:
			msg = app.commands.Unknown(&msg)
		}

		if _, err := app.bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
