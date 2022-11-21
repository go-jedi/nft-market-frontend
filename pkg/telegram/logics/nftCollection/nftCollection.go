package nftCollection

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func NftCollection(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userCollectionChoose string) error {
	if len(languageUser) > 0 {
		if len(userCollectionChoose) > 0 {
			resGetAllTokensCollection, err := requestProject.GetAllTokensCollection(userCollectionChoose)
			if err != nil {
				return err
			}
			if len(resGetAllTokensCollection) > 0 {
				var isTesting string = os.Getenv("IS_TESTING")
				var needPath string = ""
				if isTesting == "true" {
					needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
				} else {
					needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
				}
				photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s%s", needPath, "/img-need/5.jpg")))
				photo.ParseMode = "Markdown"
				if languageUser == "ru" {
					var text string = fmt.Sprintf("💠 Коллекция *%s*\n\nТокенов в коллекции: %d", resGetAllTokensCollection[0].NameCollection, resGetAllTokensCollection[0].Count)
					photo.Caption = text
					photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftCollection(resGetAllTokensCollection, "🔙 Назад")
					_, err := bot.Send(photo)
					if err != nil {
						return err
					}
				}
				if languageUser == "en" {
					var text string = fmt.Sprintf("💠 Коллекция *%s*\n\nТокенов в коллекции: %d", resGetAllTokensCollection[0].NameCollection, resGetAllTokensCollection[0].Count)
					photo.Caption = text
					photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftCollection(resGetAllTokensCollection, "🔙 Back")
					_, err := bot.Send(photo)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
