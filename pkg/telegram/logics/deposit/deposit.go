package deposit

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func Deposit(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		var isTesting string = os.Getenv("IS_TESTING")
		var needPath string = ""
		if isTesting == "true" {
			needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
		} else {
			needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
		}
		photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s%s", needPath, "/img-need/2.jpg")))
		photo.ParseMode = "Markdown"
		if languageUser == "ru" {
			photo.Caption = "Выберите метод пополнения:"
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForDeposit("🔙 Вернуться в ЛК")
		}
		if languageUser == "en" {
			photo.Caption = "Choose a method:"
			photo.ReplyMarkup = keyboard.GenKeyboardInlineForDeposit("🔙 Back to profile")
		}
		_, err := bot.Send(photo)
		if err != nil {
			return err
		}
	}

	return nil
}
