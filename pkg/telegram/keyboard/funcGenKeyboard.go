package keyboard

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GenKeyboardInlineForList(currentStep string, fullStep string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("◀️", "DG_TERMS_LIST_BACK"),
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s/%s", currentStep, fullStep), "DG_TERMS_LIST_COUNT"),
			tgbotapi.NewInlineKeyboardButtonData("▶️", "DG_TERMS_LIST_NEXT"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Согласен", "DG_TERMS_AGREE"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("◀️ Назад", "DG_WELCOME"),
		),
	)
}
