package withDrawAdmRefuse

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func WithDrawAdmRefuse(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, eventWithDraw string) error {
	resGetWithDrawEvent, err := requestProject.GetWithDrawEvent(eventWithDraw)
	if err != nil {
		return err
	}
	if len(resGetWithDrawEvent) > 0 {
		resAdminWithdrawRefuse, err := requestProject.AdminWithdrawRefuse(resGetWithDrawEvent[0].TeleId, eventWithDraw)
		if err != nil {
			return err
		}
		if resAdminWithdrawRefuse {
			resGetUserLang, err := sqlite.GetUserLang(sqliteDb, resGetWithDrawEvent[0].TeleId)
			if err != nil {
				return err
			}
			if len(resGetUserLang) > 0 {
				msg.ChatID = resGetWithDrawEvent[0].TeleId
				if resGetUserLang == "ru" {
					msg.Text = "❌ Не удалось вывести средства на указанные реквизиты, пожалуйста, свяжитесь с поддержкой."
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForWithDrawPayment("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
				}
				if resGetUserLang == "en" {
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
	}

	return nil
}
