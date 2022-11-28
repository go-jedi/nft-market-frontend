package withDrawWrite

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func WithDrawWrite(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOnListenerWatchingWithDrawWrite(sqliteDb, teleId)
		if err != nil {
			return err
		}
		if languageUser == "ru" {
			msg.Text = "Введите сумму:\n\nФормат: _1000_"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Вернуться в ЛК")
		}
		if languageUser == "en" {
			msg.Text = "Enter amount:\n\nFormat: _1000_"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Return to PA")
		}
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
