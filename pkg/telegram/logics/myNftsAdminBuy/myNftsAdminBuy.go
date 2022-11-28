package myNftsAdminBuy

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func MyNftsAdminBuy(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, uidPaymentEvent string, uidToken string) error {
	if len(languageUser) > 0 {
		resGetUserPaymentEvent, err := requestProject.GetUserPaymentEvent(uidPaymentEvent)
		if err != nil {
			return err
		}
		fmt.Println("resGetUserPaymentEvent -->", resGetUserPaymentEvent)
		if len(resGetUserPaymentEvent) > 0 {
			resAdminBuyTokenUser, err := requestProject.AdminBuyTokenUser(resGetUserPaymentEvent[0].TeleId, uidToken, resGetUserPaymentEvent[0].Price, uidPaymentEvent)
			if err != nil {
				return err
			}
			fmt.Println("resAdminBuyTokenUser -->", resAdminBuyTokenUser)
			if resAdminBuyTokenUser {
				msg.Text = "✅ Токен был успешно куплен."
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
				msg.ChatID = resGetUserPaymentEvent[0].TeleId
				if languageUser == "ru" {
					msg.Text = fmt.Sprintf("✅ Ваш NFT (%s) купили!\n%.2f $ зачислено на ваш счет", resGetUserPaymentEvent[0].NameToken, resGetUserPaymentEvent[0].Price)
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Вернуться в ЛК")
				}
				if languageUser == "en" {
					msg.Text = fmt.Sprintf("✅ Successful NFT sale (%s)", resGetUserPaymentEvent[0].NameToken)
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Return to PA")
				}
				_, err = bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
