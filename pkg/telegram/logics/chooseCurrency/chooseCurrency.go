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
				var text string = "🔸️Добро пожаловать!🔸️\n\nПеред использованием NFT площадки Rarible необходимо ознакомиться с `Рекомендациями для сообщества`"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("Рекомендации Сообщества", "✅ Подтвердить", true, languageChoose)
				msg.Text = text
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
			if languageChoose == "en" {
				msg.ParseMode = "HTML"
				var text string = "🔸️Welcome!🔸️\n\nBefore using the Rarible NFT platform, you must read the `Recommendations for the Community`"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("Community Recommendations", "✅ Accept", true, languageChoose)
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
