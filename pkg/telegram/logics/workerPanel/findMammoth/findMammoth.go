package findMammoth

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func FindMammoth(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOnListenerWatchingFindMam(sqliteDb, teleId)
		if err != nil {
			return err
		}
		msg.Text = "🦣 Введите ID мамонта, которого нужно найти\n\nФормат: _1234567890_"
		msg.ReplyMarkup = keyboard.DgFindMammothKeyboardInline
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
