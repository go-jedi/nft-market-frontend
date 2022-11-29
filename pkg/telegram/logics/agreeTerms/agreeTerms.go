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
				var text string = "🔸️Добро пожаловать!🔸️\n\nПеред использованием NFT площадки Rarible необходимо ознакомиться с `Рекомендациями для сообщества`"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("Рекомендации Сообщества", "✅ Подтвердить", true, resGetUserLang)
				msg.Text = text
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
			if resGetUserLang == "en" {
				msg.ParseMode = "HTML"
				var text string = "🔸️Welcome!🔸️\n\nBefore using the Rarible NFT platform, you must read the `Recommendations for the Community`"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("Community Recommendations", "✅ Accept", true, resGetUserLang)
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
