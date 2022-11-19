package keyboard

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GenKeyboardInlineForAgreeTerms(textBtn string, isAgreeTerms bool, lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBtn, fmt.Sprintf("DG_AREG_HOME?%t,%s", isAgreeTerms, lang)),
		),
	)
}

func GenKeyboardInlineForNftMenu(textBtn string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBtn, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForSupportMenu(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/LooksRareHelp"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForAboutMenu(textSupport string, textNews string, textReport string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/LooksRareHelp"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textNews, "https://looksrare.org/"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textReport, "https://t.me/LooksRareHelp"),
		),
	)
}

func GenKeyboardInlineForProfileMenu(textTopUp string, textWithdraw string, textMyNft string, textVerification string, textLanguaage string, languageUserChange string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textTopUp, "NM_TOP_UP"),
			tgbotapi.NewInlineKeyboardButtonData(textWithdraw, "NM_WITH_DRAW"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textMyNft, "NM_MY_NFT"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textVerification, "NM_VERIF"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textLanguaage, fmt.Sprintf("NM_CHNG_LANG?%s", languageUserChange)),
		),
	)
}

func GenKeyboardInlineForDeposit(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔸 Bitcoin", "NM_DEPOSIT_PAYMT?btc"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔹 Ethereum", "NM_DEPOSIT_PAYMT?eth"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💱 USDT", "NM_DEPOSIT_PAYMT?usdt"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForDepositPayment(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/LooksRareHelp"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForMyNfts(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForVerification(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/LooksRareHelp"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForWithDraw(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔸 Bitcoin", "NM_WITH_DRAW_PAYMT?btc"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔹 Ethereum", "NM_WITH_DRAW_PAYMT?eth"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💱 USDT", "NM_WITH_DRAW_PAYMT?usdt"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForWithDrawPayment(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/LooksRareHelp"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}
