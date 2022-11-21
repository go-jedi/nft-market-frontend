package changeMamPremium

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func ChangeMamPremium(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userChooseTeleId int64) error {
	if len(languageUser) > 0 {
		resAdminUpdatePremium, err := requestProject.AdminUpdatePremium(userChooseTeleId)
		if err != nil {
			return err
		}
		fmt.Println("resAdminUpdatePremium -->", resAdminUpdatePremium)
		if resAdminUpdatePremium {
			msg.Text = "Успешное изменение премиума пользователя"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForChangeMamPremium(userChooseTeleId)
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
