package nickPayouts

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func NickPayouts(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		resChangeVisibleName, err := requestProject.ChangeVisibleName(teleId)
		if err != nil {
			return err
		}

		if resChangeVisibleName {
			if languageUser == "ru" {
				msg.Text = "✅ Успешное изменение видимости имя"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForNickPayload("🔙 Вернуться в ЛК")
			}

			if languageUser == "en" {
				msg.Text = "✅ Successful name visibility change"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForNickPayload("🔙 Back to profile")
			}

			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
