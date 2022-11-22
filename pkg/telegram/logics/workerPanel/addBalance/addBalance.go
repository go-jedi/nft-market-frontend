package addBalance

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func AddBalance(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userTeleId int64) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOnListenerWatchingAddBalance(sqliteDb, teleId, userTeleId)
		if err != nil {
			return err
		}
		msg.Text = "🦣 Введите баланс для мамонта\n\nФормат: _1000_"
		msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(userTeleId)
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
