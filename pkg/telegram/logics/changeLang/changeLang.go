package changeLang

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/profile"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func ChangeLang(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, needLangUserChange string) error {
	if len(needLangUserChange) > 0 {
		resUpdateLanguage, err := requestProject.UpdateLanguage(teleId, needLangUserChange)
		if err != nil {
			return err
		}
		if resUpdateLanguage {
			_, err := sqliteDb.Exec("UPDATE bot_params SET lang = $1 WHERE tele_id = $2", needLangUserChange, teleId)
			if err != nil {
				return err
			}
			var text string = "✅"
			msg.Text = text
			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
			err = profile.Profile(bot, msg, teleId, userName, needLangUserChange)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
