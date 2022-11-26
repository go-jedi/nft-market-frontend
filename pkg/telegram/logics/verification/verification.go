package verification

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func Verification(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		var isTesting string = os.Getenv("IS_TESTING")
		var needPath string = ""
		if isTesting == "true" {
			needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
		} else {
			needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
		}
		photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s%s", needPath, "/img-need/4.jpg")))
		photo.ParseMode = "Markdown"
		if languageUser == "ru" {
			photo.Caption = "*Ваш аккаунт не верифицирован*\n\nДля получения инструкций по прохождению верификации напишите «Верификация» в чат технической поддержки."
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForVerification("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
		}
		if languageUser == "en" {
			photo.Caption = "*Your account is not verified*\n\nFor verification instructions, type 'Verification' in the tech support chat."
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForVerification("👨‍💻 Support", "🔙 Back to profile")
		}
		_, err := bot.Send(photo)
		if err != nil {
			return err
		}
	}

	return nil
}
