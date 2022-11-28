package withDrawAdmApprove

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func WithDrawAdmApprove(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, eventWithDraw string) error {
	if len(languageUser) > 0 {
		resGetWithDrawEvent, err := requestProject.GetWithDrawEvent(eventWithDraw)
		if err != nil {
			return err
		}
		if len(resGetWithDrawEvent) > 0 {
			resAdminWithdrawApprove, err := requestProject.AdminWithdrawApprove(resGetWithDrawEvent[0].TeleId, eventWithDraw)
			if err != nil {
				return err
			}
			fmt.Println("resAdminWithdrawApprove -->", resAdminWithdrawApprove)
			if resAdminWithdrawApprove {
				msg.ChatID = resGetWithDrawEvent[0].TeleId
				if languageUser == "ru" {
					msg.Text = fmt.Sprintf("✅ %.2f $ были успешно выведены на ваши реквизиты!", resGetWithDrawEvent[0].Price)
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Вернуться в ЛК")
				}
				if languageUser == "en" {
					msg.Text = fmt.Sprintf("✅ %.2f were successfully withdrawn to your details!", resGetWithDrawEvent[0].Price)
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Return to PA")
				}
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
				msg.ChatID = teleId
				msg.Text = "✅ Успешный вывод денежных средств."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAdminBackHome()
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
