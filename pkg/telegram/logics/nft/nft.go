package nft

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func Nft(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath("/home/dale/job/work/my-project/nft-market/frontend/img/img-nft.jpg"))
	photo.ParseMode = "Markdown"
	if languageUser == "ru" {
		photo.Caption = "💠 *Всего на маркетплейсе 13 коллекций*"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftMenu("🔙 Вернуться в ЛК")
	}
	if languageUser == "en" {
		photo.Caption = "💠 *There are 13 collections on the marketplace*"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftMenu("🔙 Back to profile")
	}
	_, err := bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}
