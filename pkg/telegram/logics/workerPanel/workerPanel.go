package workerPanel

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func WorkerPanel(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		var linkBot string = "https://t.me/nftmrkt_bot"
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = keyboard.GenKeyboardInlineForWorkerPanel()
		msg.Text = fmt.Sprintf("🔗 Ваша ссылка:\n<a href='%s?start=%d'>%s?start=%d</a>", linkBot, teleId, linkBot, teleId)
		_, err := bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
