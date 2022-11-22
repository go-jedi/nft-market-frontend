package mammothProfile

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func MammothProfile(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userChooseTeleId int64) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOffListeners(sqliteDb, teleId)
		if err != nil {
			return err
		}
		resAdminGetUserProfile, err := requestProject.AdminGetUserProfile(userChooseTeleId)
		if err != nil {
			return err
		}
		if len(resAdminGetUserProfile) > 0 {
			var isVerification string = ""
			var verificationText string = "Верифицировать"
			var isPremium string = ""
			var premiumText string = "Добавить премиум"
			if resAdminGetUserProfile[0].Verification {
				isVerification = "Да"
			} else {
				isVerification = "Нет"
				verificationText = "Верифицировать"
			}
			if resAdminGetUserProfile[0].IsPremium {
				isPremium = "Да"
			} else {
				isPremium = "Нет"
				premiumText = "Добавить премиум"
			}
			msg.ParseMode = "Markdown"
			var text string = fmt.Sprintf("👤 Профиль\n\nID: %d\nИмя: *%s*\nБаланс: *%.2f*\nНа выводе: *%.2f*\nВерифицирован: *%s*\nПремиум: *%s*",
				resAdminGetUserProfile[0].TeleId,
				resAdminGetUserProfile[0].TeleName,
				resAdminGetUserProfile[0].Balance,
				resAdminGetUserProfile[0].Conclusion,
				isVerification,
				isPremium,
			)
			msg.Text = text
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForMammothProfile(userChooseTeleId, verificationText, premiumText)
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
