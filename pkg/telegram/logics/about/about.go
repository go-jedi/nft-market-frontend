package about

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func About(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath("/home/dale/job/work/my-project/nft-market/frontend/img/img-need/3.jpg"))
	photo.ParseMode = "Markdown"
	if languageUser == "ru" {
		photo.Caption = "🔹 *О Сервисе*\n\n_LooksRare_* - торговая площадка для невзаимозаменяемых токенов (NFT). Покупайте, продавайте и открывайте для себя эксклюзивные цифровые предметы.*"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForAboutMenu("👨‍💻 Поддержка", "🗞️ Новости", "Сообщить об ошибке")
	}
	if languageUser == "en" {
		photo.Caption = "🔹 *About the Service*\n\n_LooksRare_ *is a marketplace for non-fungible tokens (NFTs). Buy, sell and discover exclusive digital items.*"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForAboutMenu("👨‍💻 Support", "🗞️ News", "Report an error")
	}
	_, err := bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}
