package changeMamPremium

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func ChangeMamPremium(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userChooseTeleId int64) error {
	if len(languageUser) > 0 {
		resAdminCheckIsPremium, err := requestProject.AdminCheckIsPremium(userChooseTeleId)
		if err != nil {
			return err
		}
		if resAdminCheckIsPremium {
			msg.Text = "Аккаунт мамонта уже является премиум"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForChangeMamPremium(userChooseTeleId)
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resAdminUpdatePremium, err := requestProject.AdminUpdatePremium(userChooseTeleId)
			if err != nil {
				return err
			}
			if resAdminUpdatePremium {
				msg.ChatID = userChooseTeleId
				if languageUser == "ru" {
					msg.Text = "✅ Ваш аккаунт получил премиум статус"
				}
				if languageUser == "en" {
					msg.Text = "✅ Your account has received a premium status"
				}
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}

				msg.ChatID = teleId
				msg.Text = "Успешное изменение премиума пользователя"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForChangeMamPremium(userChooseTeleId)
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
