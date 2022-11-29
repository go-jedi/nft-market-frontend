package changeMamVerification

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func ChangeMamVerification(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userChooseTeleId int64) error {
	if len(languageUser) > 0 {
		resAdminCheckIsVerified, err := requestProject.AdminCheckIsVerified(userChooseTeleId)
		if err != nil {
			return err
		}
		if resAdminCheckIsVerified {
			msg.Text = "Аккаунт мамонта уже является верифицированным"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForChangeMamVerification(userChooseTeleId)
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resAdminUpdateVerification, err := requestProject.AdminUpdateVerification(userChooseTeleId)
			if err != nil {
				return err
			}
			if resAdminUpdateVerification {
				msg.ChatID = userChooseTeleId
				if languageUser == "ru" {
					msg.Text = "✅ Ваш аккаунт прошёл верификацию"
				}
				if languageUser == "en" {
					msg.Text = "✅ Your account has been verified"
				}
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}

				msg.ChatID = teleId
				msg.Text = "✅ Успешное изменение верификации пользователя"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForChangeMamVerification(userChooseTeleId)
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
