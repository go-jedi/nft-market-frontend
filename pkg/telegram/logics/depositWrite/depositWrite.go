package depositWrite

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func DepositWrite(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userPaymentChoose string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOnListenerWatchingWritePrice(sqliteDb, teleId, userPaymentChoose)
		if err != nil {
			return err
		}
		if languageUser == "ru" {
			msg.Text = "Введите сумму пополнения\n\nФормат: _1000_"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Назад в профиль")
		}
		if languageUser == "en" {
			msg.Text = "Enter the replenishment amount\n\nFormat: _1000_"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Back to profile")
		}
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
