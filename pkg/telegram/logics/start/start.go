package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetStart(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string) error {
	var text string = "Вызов функции по /start команде"
	msg.Text = text
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
