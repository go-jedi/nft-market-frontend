package nftTokenBuy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func NftTokenBuy(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userTokenChoose string) error {
	if len(languageUser) > 0 {
		if len(userTokenChoose) > 0 {
			if languageUser == "ru" {
				msg.Text = "❌ На вашем счете недостаточно средств для покупки токена."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuy("🔙 Назад", userTokenChoose)
			}
			if languageUser == "en" {
				msg.Text = "❌ You don't have enough funds to buy this token."
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuy("🔙 Back", userTokenChoose)
			}
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
