package about

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func About(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	var isTesting string = os.Getenv("IS_TESTING")
	var needPath string = ""
	if isTesting == "true" {
		needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
	} else {
		needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
	}
	photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s%s", needPath, "/img-need/3.jpg")))
	photo.ParseMode = "Markdown"
	if languageUser == "ru" {
		photo.Caption = "🔹 *О Сервисе*\n\n_Rarible_* - торговая площадка для невзаимозаменяемых токенов (NFT). Покупайте, продавайте и открывайте для себя эксклюзивные цифровые предметы.*"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForAboutMenu("📄 Наш сайт", "Политика конфиденциальности", " Условия использования", "Сообщить об ошибке")
	}
	if languageUser == "en" {
		photo.Caption = "🔹 *About the Service*\n\n_Rarible_ *is a marketplace for non-fungible tokens (NFTs). Buy, sell and discover exclusive digital items.*"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForAboutMenu("📄 Our website", "Privacy Policy", "Terms of Use", "Report an error")
	}
	_, err := bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}
