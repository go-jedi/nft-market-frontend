package myMammoths

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func modifyDate(needTime string) (string, error) {
	var needDateFormat string = ""
	needTimeSplit := strings.Split(needTime, "T")
	needTimeDateSplit := strings.Split(needTimeSplit[0], "-")
	needTimeTimeSplit := strings.Split(needTimeSplit[1], ".")
	needDateFormat = fmt.Sprintf("%s.%s.%s %s", needTimeDateSplit[2], needTimeDateSplit[1], needTimeDateSplit[0], needTimeTimeSplit[0])
	return needDateFormat, nil
}

func MyMammoths(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	if len(languageUser) > 0 {
		resGetUserReferral, err := requestProject.GetUserReferral(teleId)
		if err != nil {
			return err
		}
		if len(resGetUserReferral) > 0 {
			msg.ParseMode = "HTML"
			var text string = fmt.Sprintf("🦣 Мои мамонты\n\nВсего: <b>%d</b>\n\n", len(resGetUserReferral))
			for _, value := range resGetUserReferral {
				resModifyDate, err := modifyDate(value.Created)
				if err != nil {
					return err
				}
				text += fmt.Sprintf("🦣 %s\n📅 %s\n📂 <a href='%d'>/u%d</a>\n\n", value.TeleName, resModifyDate, value.TeleId, value.TeleId)
			}
			msg.Text = text
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForMyMammoths()
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			msg.ParseMode = "HTML"
			var text string = fmt.Sprintf("🦣 Мои мамонты\n\nВсего: <b>%d</b>", len(resGetUserReferral))
			msg.Text = text
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForMyMammoths()
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
