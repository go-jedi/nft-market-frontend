package workerPanel

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func WorkerPanel(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOffListeners(sqliteDb, teleId)
		if err != nil {
			return err
		}
		var linkBot string = "https://t.me/nftmrkttest_bot"
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = keyboard.GenKeyboardInlineForWorkerPanel()
		msg.Text = fmt.Sprintf("🔗 Ваша ссылка:\n<a href='%s?start=%d'>%s?start=%d</a>", linkBot, teleId, linkBot, teleId)
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
