package agreeTerms

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func AgreeTerms(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, currencyChoose string) error {
	fmt.Println("currencyChoose -->", currencyChoose)
	if len(currencyChoose) > 0 {
		resUpdateCurrency, err := requestProject.UpdateCurrency(teleId, currencyChoose)
		if err != nil {
			return err
		}
		if resUpdateCurrency {
			resGetUserLang, err := sqlite.GetUserLang(sqliteDb, teleId)
			if err != nil {
				return err
			}
			fmt.Println("resGetUserLang -->", resGetUserLang)
			if resGetUserLang == "ru" {
				msg.ParseMode = "HTML"
				var text string = "Подтвердите, что вы не бот.\n\nНажмимая “Подтвердить“, Вы принимаете условия <a href='https://static.rarible.com/terms.pdf'>пользовательского соглашения</a>, <a href='https://static.rarible.com/privacy.pdf'>Условия конфиденциальности</a>."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("✅ Подтвердить", true, resGetUserLang)
				msg.Text = text
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
			if resGetUserLang == "en" {
				msg.ParseMode = "HTML"
				var text string = "Please, confirm you're not a robot.\n\nBy pressing “Accept“ you confirm that you've read and accept our <a href='https://static.rarible.com/terms.pdf'>Terms</a>, <a href='https://static.rarible.com/privacy.pdf'>Privacy</a>."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("✅ Accept", true, resGetUserLang)
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
