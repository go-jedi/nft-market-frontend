package nft

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func Nft(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		resGetAllCollections, err := requestProject.GetAllCollections()
		if err != nil {
			return err
		}
		if len(resGetAllCollections) > 0 {
			photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath("/home/dale/job/work/my-project/nft-market/frontend/img/img-need/1.jpg"))
			photo.ParseMode = "Markdown"
			if languageUser == "ru" {
				var nameCollection string = ""
				if len(resGetAllCollections) == 1 {
					nameCollection = "коллекция"
				}
				if len(resGetAllCollections) > 1 {
					nameCollection = "коллекций"
				}
				photo.Caption = fmt.Sprintf("💠 *Всего на маркетплейсе %d %s*", len(resGetAllCollections), nameCollection)
				photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftMenu(resGetAllCollections, "🔙 Вернуться в ЛК")
			}
			if languageUser == "en" {
				var nameCollection string = ""
				if len(resGetAllCollections) == 1 {
					nameCollection = "collection"
				}
				if len(resGetAllCollections) > 1 {
					nameCollection = "collections"
				}
				photo.Caption = fmt.Sprintf("💠 *There are %d %s on the marketplace*", len(resGetAllCollections), nameCollection)
				photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftMenu(resGetAllCollections, "🔙 Back to profile")
			}
			_, err := bot.Send(photo)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
