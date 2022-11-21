package chooseCurrency

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func ChooseCurrency(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageChoose string) error {
	if len(languageChoose) > 0 {
		resUpdateLanguage, err := requestProject.UpdateLanguage(teleId, languageChoose)
		if err != nil {
			return err
		}
		if resUpdateLanguage {
			_, err = sqliteDb.Exec("INSERT INTO bot_params(tele_id, lang) VALUES($1, $2)", teleId, languageChoose)
			if err != nil {
				return err
			}
			if languageChoose == "ru" {
				msg.ParseMode = "HTML"
				var text string = "Подтвердите, что вы не бот.\n\nНажмимая “Подтвердить“, Вы принимаете условия <a href='https://static.rarible.com/terms.pdf'>пользовательского соглашения</a>, <a href='https://static.rarible.com/privacy.pdf'>Условия конфиденциальности</a>."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("✅ Подтвердить", true, languageChoose)
				msg.Text = text
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
			if languageChoose == "en" {
				msg.ParseMode = "HTML"
				var text string = "Please, confirm you're not a robot.\n\nBy pressing “Accept“ you confirm that you've read and accept our <a href='https://static.rarible.com/terms.pdf'>Terms</a>, <a href='https://static.rarible.com/privacy.pdf'>Privacy</a>."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("✅ Accept", true, languageChoose)
				msg.Text = text
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
