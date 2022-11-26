package workerPanel

import (
	"database/sql"
	"fmt"
	"os"

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
		var linkBot string = os.Getenv("LINK_BOT")
		msg.ParseMode = "Markdown"
		msg.ReplyMarkup = keyboard.GenKeyboardInlineForWorkerPanel()
		msg.Text = fmt.Sprintf("🔗 Твоя ссылка: `%s?start=%d`", linkBot, teleId)
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
