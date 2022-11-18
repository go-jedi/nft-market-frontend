package terms

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetTerms(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string) error {
	var text string = "Вызов функции по обработки inlineButton"
	msg.Text = text
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
