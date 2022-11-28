package myNftsToken

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func MyNftsToken(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userTokenChoose string, isSellToken string) error {
	if len(languageUser) > 0 {
		if len(userTokenChoose) > 0 {
			resGetToken, err := requestProject.GetToken(userTokenChoose)
			if err != nil {
				return err
			}
			if len(resGetToken) > 0 {
				var isTesting string = os.Getenv("IS_TESTING")
				var needPath string = ""
				if isTesting == "true" {
					needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
				} else {
					needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
				}
				photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s/nft/%s.jpg", needPath, resGetToken[0].TokenUid)))
				photo.ParseMode = "Markdown"
				if languageUser == "ru" {
					photo.Caption = fmt.Sprintf("💠 Токен *%s*\n\n🗂 Коллекция: *%s*\n👩‍💻 Автор: *%s*\n🔹 Блокчейн: *%s*\n\n💸 Цена: *$%.2f*",
						resGetToken[0].Name,
						resGetToken[0].NameCollection,
						resGetToken[0].Author,
						resGetToken[0].Blockchain,
						resGetToken[0].Price,
					)
					if isSellToken == "yes" {
						photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftsTokenSell("🔙 Назад")
					} else {
						photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftsToken(resGetToken, "💰 Продать", "🔙 Назад")
					}
				}
				if languageUser == "en" {
					photo.Caption = fmt.Sprintf("💠 Token *%s*\n\n🗂 Collection: *%s*\n👩‍💻 Author: *%s*\n🔹 Blockchain: *%s*\n\n💸 Price: *$%.2f*",
						resGetToken[0].Name,
						resGetToken[0].NameCollection,
						resGetToken[0].Author,
						resGetToken[0].Blockchain,
						resGetToken[0].Price,
					)
					if isSellToken == "yes" {
						photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftsTokenSell("🔙 Back")
					} else {
						photo.ReplyMarkup = keyboard.GenKeyboardInlineForNftsToken(resGetToken, "💰 Sell", "🔙 Back")
					}
				}
				_, err := bot.Send(photo)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
