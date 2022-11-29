package withDrawAdmApprove

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func WithDrawAdmApprove(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, eventWithDraw string) error {
	resGetWithDrawEvent, err := requestProject.GetWithDrawEvent(eventWithDraw)
	if err != nil {
		return err
	}
	if len(resGetWithDrawEvent) > 0 {
		resAdminWithdrawApprove, err := requestProject.AdminWithdrawApprove(resGetWithDrawEvent[0].TeleId, eventWithDraw)
		if err != nil {
			return err
		}
		if resAdminWithdrawApprove {
			resGetUserLang, err := sqlite.GetUserLang(sqliteDb, resGetWithDrawEvent[0].TeleId)
			if err != nil {
				return err
			}
			if len(resGetUserLang) > 0 {
				msg.ChatID = resGetWithDrawEvent[0].TeleId
				if resGetUserLang == "ru" {
					msg.Text = fmt.Sprintf("✅ %.2f $ были успешно выведены на ваши реквизиты!", resGetWithDrawEvent[0].Price)
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositWrite("🔙 Вернуться в ЛК")
				}
				if resGetUserLang == "en" {
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
