package myMammoths

import (
	"database/sql"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func MyMammoths(bot *tgbotapi.BotAPI, sqliteDb *sql.DB, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		err := sqlite.TurnOffListeners(sqliteDb, teleId)
		if err != nil {
			return err
		}
		resCheckUserReferral, err := requestProject.CheckUserReferral(teleId)
		if err != nil {
			return err
		}
		if len(resCheckUserReferral) > 0 {
			msg.ParseMode = "HTML"
			var text string = fmt.Sprintf("🦣 Мои мамонты\n\nВсего: <b>%d</b>", resCheckUserReferral[0].Count)
			msg.Text = text
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForMyMammoths()
			_, err = bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
