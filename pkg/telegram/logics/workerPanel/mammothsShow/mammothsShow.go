package mammothsShow

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/helperFunc"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func MammothsShow(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, limitShowMammoths int) error {
	if len(languageUser) > 0 {
		resGetUserReferral, err := requestProject.GetUsersReferral(teleId, limitShowMammoths)
		if err != nil {
			return err
		}
		if len(resGetUserReferral) > 0 {
			var needTitle string = ""
			if limitShowMammoths == 5 {
				needTitle = "5️⃣ Последние 5"
			}
			if limitShowMammoths == 50 {
				needTitle = "5️⃣0️⃣ Последние 50"
			}
			msg.ParseMode = "HTML"
			var text string = fmt.Sprintf("%s\n\n", needTitle)
			for _, value := range resGetUserReferral {
				resModifyDate, err := helperFunc.ModifyDate(value.Created)
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
		}
	}
	return nil
}
