package deposit

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func Deposit(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath("/home/dale/job/work/my-project/nft-market/frontend/img/img-deposit.jpg"))
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
