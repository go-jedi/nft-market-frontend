package addMammoth

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func AddMammoth(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOnListenerWatchingAddMam(sqliteDb, teleId)
		if err != nil {
			return err
		}
		msg.Text = "🦣 Введите ID мамонта и его имя, которого нужно добавить\n\nФормат: _1234567890 Джони_"
		msg.ReplyMarkup = keyboard.DgAddMammothKeyboardInline
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
