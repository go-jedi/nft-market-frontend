package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var DgLangKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Русский", "NM_LANGUAGE?ru"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇸 English", "NM_LANGUAGE?en"),
	),
)

var DgCurrencyKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺 RUB", "NM_CURRENCY?rub"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇦 UAH", "NM_CURRENCY?uah"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇺🇸 USD", "NM_CURRENCY?usd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇪🇺 EUR", "NM_CURRENCY?eur"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇵🇱 PLN", "NM_CURRENCY?pln"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇧🇾 BYN", "NM_CURRENCY?byn"),
	),
)

var DgAddMammothKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 Отменить добавление", "NM_WORKPANEL_MAM"),
	),
)

var DgAddMammothSuccessKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 К мамонтам", "NM_WORKPANEL_MAM"),
	),
)

var DgFindMammothKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 Отменить поиск", "NM_WORKPANEL"),
	),
)

var DgFindMammothSuccessKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
	),
)

var DgChangeMamMinimalKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
	),
)
