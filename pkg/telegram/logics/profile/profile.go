package profile

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func modifyDate(needTime time.Time) (string, error) {
	var needDateFormat string = ""
	needTimeSplit := strings.Split(needTime.String(), " ")
	needTimeDateSplit := strings.Split(needTimeSplit[0], "-")
	needTimeTimeSplit := strings.Split(needTimeSplit[1], ".")
	needDateFormat = fmt.Sprintf("%s.%s.%s %s", needTimeDateSplit[2], needTimeDateSplit[1], needTimeDateSplit[0], needTimeTimeSplit[0])
	return needDateFormat, nil
}

func Profile(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	resGetUserProfile, err := requestProject.GetUserProfile(teleId)
	if err != nil {
		return err
	}
	fmt.Println("resGetUserProfile -->", resGetUserProfile)
	if len(resGetUserProfile) > 0 {
		photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath("/home/dale/job/work/my-project/nft-market/frontend/img/img-profile.jpg"))
		photo.ParseMode = "Markdown"
		loc, err := time.LoadLocation("Europe/Moscow")
		now := time.Now().In(loc)
		if err != nil {
			return err
		}
		resModifyDate, err := modifyDate(now)
		if err != nil {
			return err
		}
		if languageUser == "ru" {
			var isVerification string
			if resGetUserProfile[0].Verification {
				isVerification = "✅ *Верифицирован*"
			} else {
				isVerification = "⚠️ *Не верифицирован*"
			}
			photo.Caption = fmt.Sprintf("*Личный кабинет*\n\nБаланс: *%d $*\nНа выводе: *%d $*\n\nВерификация: %s\nВаш ID: [%d](tg://user?id=%d)\n\nДата и время: %s",
				resGetUserProfile[0].Balance,
				resGetUserProfile[0].Conclusion,
				isVerification,
				teleId,
				teleId,
				resModifyDate,
			)
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForProfileMenu("📥 Пополнить", "📤 Вывести", "🖼 Мои NFT", "📝 Верификация", "🇺🇸 English language", "en")
		}

		if languageUser == "en" {
			var isVerification string
			if resGetUserProfile[0].Verification {
				isVerification = "✅ *Verified*"
			} else {
				isVerification = "⚠️ *Not verified*"
			}
			photo.Caption = fmt.Sprintf("*Personal account*\n\nBalance: *%d $*\nWithdrawal: *%d $*\n\nVerification: %s\nYour ID: [%d](tg://user?id=%d)\n\nDate and time: %s",
				resGetUserProfile[0].Balance,
				resGetUserProfile[0].Conclusion,
				isVerification,
				teleId,
				teleId,
				resModifyDate,
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
