package mammothProfile

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func MammothProfile(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userChooseTeleId int64) error {
	if len(languageUser) > 0 {
		resAdminGetUserProfile, err := requestProject.AdminGetUserProfile(userChooseTeleId)
		if err != nil {
			return err
		}
		if len(resAdminGetUserProfile) > 0 {
			var isVerification string = ""
			var isPremium string = ""
			var premiumText string = ""
			if resAdminGetUserProfile[0].Verification {
				isVerification = "Да"
			} else {
				isVerification = "Нет"
			}
			if resAdminGetUserProfile[0].IsPremium {
				isPremium = "Да"
				premiumText = "Убрать премиум"
			} else {
				isPremium = "Нет"
				premiumText = "Добавить премиум"
			}
			msg.ParseMode = "Markdown"
			var text string = fmt.Sprintf("👤 Профиль\n\nID: %d\nИмя: *%s*\nБаланс: *%d*\nНа выводе: *%d*\nВерифицирован: *%s*\nПремиум: *%s*",
				resAdminGetUserProfile[0].TeleId,
				resAdminGetUserProfile[0].TeleName,
				resAdminGetUserProfile[0].Balance,
				resAdminGetUserProfile[0].Conclusion,
				isVerification,
				isPremium,
			)
			msg.Text = text
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForMammothProfile(userChooseTeleId, premiumText)
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
