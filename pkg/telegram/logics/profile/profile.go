package profile

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func Profile(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	resGetUserProfile, err := requestProject.GetUserProfile(teleId)
	if err != nil {
		return err
	}
	fmt.Println("resGetUserProfile -->", resGetUserProfile)
	if len(resGetUserProfile) > 0 {
		var isTesting string = os.Getenv("IS_TESTING")
		var needPath string = ""
		if isTesting == "true" {
			needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
		} else {
			needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
		}
		photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s%s", needPath, "/img-need/4.jpg")))
		photo.ParseMode = "Markdown"
		if err != nil {
			return err
		}
		if languageUser == "ru" {
			var isVerification string
			var isPremium string
			if resGetUserProfile[0].Verification {
				isVerification = "✅ *Верифицирован*"
			} else {
				isVerification = "⚠️ *Не верифицирован*"
			}
			if resGetUserProfile[0].IsPremium {
				isPremium = "✅ *Премиум*"
			} else {
				isPremium = "⭕️ *Не премиум*"
			}
			photo.Caption = fmt.Sprintf("*Личный кабинет*\n\nБаланс: *%d $*\nНа выводе: *%d $*\n\nВерификация: %s\nСтатус аккаунта: %s\nВаш ID: [%d](tg://user?id=%d)",
				resGetUserProfile[0].Balance,
				resGetUserProfile[0].Conclusion,
				isVerification,
				isPremium,
				teleId,
				teleId,
			)
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForProfileMenu("📥 Пополнить", "📤 Вывести", "🖼 Мои NFT", "📝 Верификация", "🇺🇸 English language", "en")
		}

		if languageUser == "en" {
			var isVerification string
			var isPremium string
			if resGetUserProfile[0].Verification {
				isVerification = "✅ *Verified*"
			} else {
				isVerification = "⚠️ *Not verified*"
			}
			if resGetUserProfile[0].IsPremium {
				isPremium = "✅ *Premium*"
			} else {
				isPremium = "⭕️ *Not premium*"
			}
			photo.Caption = fmt.Sprintf("*Personal account*\n\nBalance: *%d $*\nWithdrawal: *%d $*\n\nVerification: %s\nStatus Account: %s\nYour ID: [%d](tg://user?id=%d)",
				resGetUserProfile[0].Balance,
				resGetUserProfile[0].Conclusion,
				isVerification,
				isPremium,
				teleId,
				teleId,
			)
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForProfileMenu("📥 Deposit", "📤 Withdraw", "🖼 My NFTs", "📝 Verification", "🇷🇺 Русский язык", "ru")
		}

		_, err = bot.Send(photo)
		if err != nil {
			return err
		}
	}

	return nil
}
