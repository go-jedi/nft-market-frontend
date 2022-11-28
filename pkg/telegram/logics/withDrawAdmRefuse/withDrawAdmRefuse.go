package withDrawAdmRefuse

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func WithDrawAdmRefuse(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, eventWithDraw string) error {
	if len(languageUser) > 0 {
		resGetWithDrawEvent, err := requestProject.GetWithDrawEvent(eventWithDraw)
		if err != nil {
			return err
		}
		if len(resGetWithDrawEvent) > 0 {
			msg.ChatID = resGetWithDrawEvent[0].TeleId
			if languageUser == "ru" {
				msg.Text = "❌ Не удалось вывести средства на указанные реквизиты, пожалуйста, свяжитесь с поддержкой."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForWithDrawPayment("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
			}
			if languageUser == "en" {
				msg.Text = "❌ Failed to withdraw funds to the specified details, please contact support."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForWithDrawPayment("👨‍💻 Support", "🔙 Back to profile")
			}
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
			msg.ChatID = teleId
			msg.Text = "✅ Успешный отказ мамонту на вывод средств."
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForAdminBackHome()
			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
