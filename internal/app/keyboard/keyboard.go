package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Keyboard struct {
}

func New() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) StartThinking() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Начать размышление", "startThinking"),
		),
	)
}
