package blockUser

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func BlockUser(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userTeleId int64) error {
	if len(languageUser) > 0 {
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(userTeleId)
		if err != nil {
			return err
		}
		fmt.Println("resCheckIsBlockUser -->", resCheckIsBlockUser)
		if resCheckIsBlockUser {
			msg.Text = "Мамонт уже является заблокированным"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(userTeleId)
			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resBlockUser, err := requestProject.BlockUser(userTeleId)
			if err != nil {
				return err
			}
			if resBlockUser {
				msg.Text = "Мамонт успешно заблокирован"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(userTeleId)
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
