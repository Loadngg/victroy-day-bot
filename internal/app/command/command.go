package command

import (
	"victory-day-bot/internal/app/keyboard"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command struct {
	keyboard *keyboard.Keyboard
}

func New() *Command {
	c := &Command{}

	c.keyboard = keyboard.New()

	return c
}

func (c *Command) Start(msg *tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "Начинаем квест"
	msg.ReplyMarkup = c.keyboard.StartThinking()

	return *msg
}

func (c *Command) Unknown(msg *tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "Я не знаю данную команду"

	return *msg
}
