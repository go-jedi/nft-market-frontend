package keyboard

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func GenKeyboardInlineForAgreeTerms(textBtn string, isAgreeTerms bool, lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBtn, fmt.Sprintf("DG_AREG_HOME?%t,%s", isAgreeTerms, lang)),
		),
	)
}

func GenKeyboardInlineForNftMenu(collections []requestProject.Collection, textBtn string) tgbotapi.InlineKeyboardMarkup {
	var keyboardCollection tgbotapi.InlineKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup()
	if len(collections) > 0 {
		for _, value := range collections {
			keyboardCollection.InlineKeyboard = append(keyboardCollection.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s (%d)", value.Name, value.Count), fmt.Sprintf("NM_NFT_COLL?%s", value.CollectionUid)),
			))
		}
	}
	keyboardCollection.InlineKeyboard = append(keyboardCollection.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(textBtn, "NM_PROFILE"),
	))

	return keyboardCollection
}

func GenKeyboardInlineForSupportMenu(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/mySupportWorkBot"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForAboutMenu(textSupport string, textNews string, textReport string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/mySupportWorkBot"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textNews, "https://looksrare.org/"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textReport, "https://t.me/mySupportWorkBot"),
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
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/mySupportWorkBot"),
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
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/mySupportWorkBot"),
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
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/mySupportWorkBot"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForNftCollection(tokensCollection []requestProject.TokensGetByCollection, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	var keyboardtokensCollection tgbotapi.InlineKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup()
	if len(tokensCollection) > 0 {
		for _, value := range tokensCollection {
			keyboardtokensCollection.InlineKeyboard = append(keyboardtokensCollection.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s ($%.2f)", value.NameToken, value.PriceToken), fmt.Sprintf("NM_NFT_COLL_T?%s", value.TokenUid)),
			))
		}
	}
	keyboardtokensCollection.InlineKeyboard = append(keyboardtokensCollection.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_NFT"),
	))

	return keyboardtokensCollection
}

func GenKeyboardInlineForNftToken(token []requestProject.Token, textBuy string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBuy, fmt.Sprintf("NM_NFT_COLL_TB?%s", token[0].TokenUid)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, fmt.Sprintf("NM_NFT_COLL?%s", token[0].UidCollection)),
		),
	)
}

func GenKeyboardInlineForNftTokenBuy(textBackProfile string, tokenUid string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, fmt.Sprintf("NM_NFT_COLL_T?%s", tokenUid)),
		),
	)
}

func GenKeyboardInlineForWorkerPanel() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🦣 Мои мамонты", "NM_WORKPANEL_MAM"),
		),
	)
}

func GenKeyboardInlineForMyMammoths() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
		),
	)
}

func GenKeyboardInlineForMammothProfile(teleId int64, textPremium string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textPremium, fmt.Sprintf("NM_WORKPANEL_MAM_PREM?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
		),
	)
}

func GenKeyboardInlineForChangeMamPremium(teleId int64) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться к пользователю", fmt.Sprintf("NM_WORKPANEL_MAM_US?%d", teleId)),
		),
	)
}
