package telegram

import (
	"fmt"

	"github.com/Vladosya/nft-market-frontend/pkg/telegram/logics/start"
	"github.com/Vladosya/nft-market-frontend/pkg/telegram/logics/terms"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.From.ID, "Я не знаю такой команды")

	switch message.Command() {
	case CommandStart:
		msg.ParseMode = "Markdown"
		err := start.GetStart(b.Bot, msg, message.From.ID, message.From.UserName)
		if err != nil {
			return err
		}
	case CommandHelp:
		msg.Text = "Ты ввёл команду /help"
		// msg.ReplyMarkup = numericKeyboard
		_, err := b.Bot.Send(msg)
		if err != nil {
			return err
		}
	default:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
		_, err := b.Bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) callbackQuery(callbackQuery tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "callbackQuery")
	switch callbackQuery.Data {
	case "DG_TERMS":
		fmt.Println("DG_TERMSDG_TERMS")
		terms.GetTerms(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName)
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.Bot.Send(msg)
}
