package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
// 		tgbotapi.NewInlineKeyboardButtonData("2", "22"),
// 		tgbotapi.NewInlineKeyboardButtonData("3", "33"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("4", "44"),
// 		tgbotapi.NewInlineKeyboardButtonData("5", "55"),
// 		tgbotapi.NewInlineKeyboardButtonData("6", "66"),
// 	),
// )

var DgTermsKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("📋 Пользовательское соглашение​", "DG_TERMS"),
	),
)

var DgHomeKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🏠​ Меню", "DG_HOME"),
	),
)

var DgTermsQtnOneKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("📄 Анкета", "DG_TERMS_QTN_ONE"),
	),
)
