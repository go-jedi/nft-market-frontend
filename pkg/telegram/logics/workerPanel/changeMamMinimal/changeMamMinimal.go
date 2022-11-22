package changeMamMinimal

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func ChangeMamMinimal(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOnListenerWatchingChangeMinLink(sqliteDb, teleId)
		if err != nil {
			return err
		}
		msg.Text = "Введите минималку, которая будет устанавливаться мамонтам с твоей рефкой:"
		msg.ReplyMarkup = keyboard.DgChangeMamMinimalKeyboardInline
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
