package homeAfterReg

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func HomeAfterReg(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, isAgreeTerms bool, languageUser string) error {
	fmt.Println("isAgreeTerms -->", isAgreeTerms)
	if isAgreeTerms {
		resAgreeTerms, err := requestProject.AgreeTerms(teleId)
		if err != nil {
			return err
		}
		if resAgreeTerms {
			var text string = "✅"
			msg.Text = text
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
			var textTwo string = ""
			if languageUser == "ru" {
				textTwo = "Если у вас не появилось меню, то напишите /start"
			}
			if languageUser == "en" {
				textTwo = "In case the menu did not appear, send /start or press it"
			}
			msg.Text = textTwo
			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
			resCheckIsAdmin, err := requestProject.CheckIsAdmin(teleId)
			if err != nil {
				return err
			}
			if !resCheckIsAdmin {
				resGetAdminByUser, err := requestProject.GetAdminByUser(teleId)
				if err != nil {
					return err
				}
				if len(resGetAdminByUser) > 0 {
					msg.ChatID = resGetAdminByUser[0].TeleId
					msg.ParseMode = "HTML"
					msg.Text = fmt.Sprintf("✅ Мамонт %s @%s /u%d) <b>зарегистрировался по твоей ссылке</b>", userName, userName, teleId)
					_, err := bot.Send(msg)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
