package myNftsTokenSell

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func MyNftsTokenSell(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userTokenChoose string) error {
	if len(languageUser) > 0 {
		if len(userTokenChoose) > 0 {
			err := sqlite.TurnOnListenerWatchingNftSell(sqliteDb, teleId, userTokenChoose)
			if err != nil {
				return err
			}
			if languageUser == "ru" {
				msg.Text = "Введите цену в долларах, за которую хотите продать данный NFT\n\nФормат: _1000_"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForTokenSell("🔙 Назад")
			}
			if languageUser == "en" {
				msg.Text = "Enter the dollar price for which I want to sell this NFT\n\nFormat: _1000_"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForTokenSell("🔙 Back")
			}
			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
