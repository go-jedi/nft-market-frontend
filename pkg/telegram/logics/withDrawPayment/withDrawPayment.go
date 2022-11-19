package withDrawPayment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func WithDrawPayment(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userPaymentChoose string) error {
	if len(languageUser) > 0 {
		if len(userPaymentChoose) > 0 {
			if languageUser == "ru" {
				msg.Text = "Withdrawal in crypto\n\n⚠️ Dear user, due to technical reasons crypto withdrawal is only available through our support team."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForWithDrawPayment("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}

			if languageUser == "en" {
				msg.Text = "Withdrawal in crypto\n\n⚠️ Dear user, due to technical reasons crypto withdrawal is only available through our support team."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForWithDrawPayment("👨‍💻 Support", "🔙 Back to profile")
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
