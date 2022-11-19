package telegram

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot      *tgbotapi.BotAPI
	SqliteDb *sql.DB
}

func NewBot(bot *tgbotapi.BotAPI, sqliteDb *sql.DB) *Bot {
	return &Bot{Bot: bot, SqliteDb: sqliteDb}
}

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return nil
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			b.callbackQuery(*update.CallbackQuery)
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.Bot.GetUpdatesChan(u), nil
}
