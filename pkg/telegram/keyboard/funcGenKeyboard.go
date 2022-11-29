package keyboard

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func GenKeyboardInlineForAgreeTerms(textReccomendation string, textBtn string, isAgreeTerms bool, lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textReccomendation, "https://rarible.com/community-guidelines"),
		),
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
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/Rarible_Support_bot"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForAboutMenu(textSupport string, textNews string, termsOfUse string, textReport string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://rarible.com/"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textNews, "https://static.rarible.com/privacy.pdf"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(termsOfUse, "https://static.rarible.com/terms.pdf"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textReport, "https://t.me/Rarible_Support_bot"),
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

func GenKeyboardInlineForProfileMenuAdmin(textTopUp string, textWithdraw string, textMyNft string, textVerification string, textLanguaage string, languageUserChange string, hideText string) tgbotapi.InlineKeyboardMarkup {
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
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(hideText, fmt.Sprintf("NM_HIDE_NCK?%s", languageUserChange)),
		),
	)
}

func GenKeyboardInlineForDeposit(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔸 Bitcoin", "NM_DEPOSIT_WRT?btc"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔹 Ethereum", "NM_DEPOSIT_WRT?eth"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💱 USDT", "NM_DEPOSIT_WRT?usdt"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForDepositPayment(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/Rarible_Support_bot"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForNickPayload(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForVerification(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/Rarible_Support_bot"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

// func GenKeyboardInlineForWithDraw(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
// 	return tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("🔸 Bitcoin", "NM_WITH_DRAW_PAYMT?btc"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("🔹 Ethereum", "NM_WITH_DRAW_PAYMT?eth"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("💱 USDT", "NM_WITH_DRAW_PAYMT?usdt"),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
// 		),
// 	)
// }

func GenKeyboardInlineForWithDraw(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔸 Bitcoin", "NM_WITH_DRAW_WR"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔹 Ethereum", "NM_WITH_DRAW_WR"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💱 USDT", "NM_WITH_DRAW_WR"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForWithDrawPayment(textSupport string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(textSupport, "https://t.me/Rarible_Support_bot"),
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

func GenKeyboardInlineForNftTokenBuyHaveToken(token []requestProject.Token, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, fmt.Sprintf("NM_NFT_COLL?%s", token[0].UidCollection)),
		),
	)
}

func GenKeyboardInlineForWorkerPanel() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🦣 Мои мамонты", "NM_WORKPANEL_MAM"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔍 Поиск", "NM_WORKPANEL_SRC"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Изменить минималку", "NM_WORKPANEL_CMIN"),
		),
	)
}

func GenKeyboardInlineForMyMammoths() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("5️⃣ Последние 5", "NM_WORKPANEL_SF?5"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("5️⃣0️⃣ Последние 50", "NM_WORKPANEL_SF?50"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Добавить", "NM_WORKPANEL_AD"),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// tgbotapi.NewInlineKeyboardButtonData("💥 Удалить всех мамонтов", "NM_WORKPANEL_DAM"),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
		),
	)
}

func GenKeyboardInlineForAdminBackHome() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
		),
	)
}

func GenKeyboardInlineForMammothProfile(teleId int64, verificationText string, textPremium string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(verificationText, fmt.Sprintf("NM_WORKPANEL_MAM_VERIF?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textPremium, fmt.Sprintf("NM_WORKPANEL_MAM_PREM?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💰 Пополнить баланс", fmt.Sprintf("NM_WORKPANEL_ADB?%d", teleId)),
			tgbotapi.NewInlineKeyboardButtonData("📥 Минималка", fmt.Sprintf("NM_WORKPANEL_MNU?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💸 Изменить баланс", fmt.Sprintf("NM_WORKPANEL_CHB?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✉️ Сообщение мамонту", fmt.Sprintf("NM_WORKPANEL_MSM?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("♻️ Обновить", fmt.Sprintf("NM_WORKPANEL_MAM_US?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❌ Заблокировать мамонта", fmt.Sprintf("NM_WORKPANEL_BUS?%d", teleId)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 В меню", "NM_WORKPANEL"),
			// tgbotapi.NewInlineKeyboardButtonData("🔍 Искать еще", "NM_WORKPANEL"),
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

func GenKeyboardInlineForChangeMamVerification(teleId int64) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться к пользователю", fmt.Sprintf("NM_WORKPANEL_MAM_US?%d", teleId)),
		),
	)
}

func GenKeyboardInlineForAddBalance(teleId int64) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Назад в профиль", fmt.Sprintf("NM_WORKPANEL_MAM_US?%d", teleId)),
		),
	)
}

func GenKeyboardInlineForDepositWrite(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
		),
	)
}

func GenKeyboardInlineForMyNfts(userNft requestProject.UserGetNft, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	var keyboardUserNft tgbotapi.InlineKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup()
	if len(userNft.NftSell) > 0 {
		for _, value := range userNft.NftSell {
			keyboardUserNft.InlineKeyboard = append(keyboardUserNft.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("◾️ | %s", value.Name), fmt.Sprintf("NM_MY_NFT_N?%s,yes", value.TokenUid)),
			))
		}
	}
	if len(userNft.NftBuy) > 0 {
		for _, value := range userNft.NftBuy {
			keyboardUserNft.InlineKeyboard = append(keyboardUserNft.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🔹 | %s ($%.2f)", value.Name, value.Price), fmt.Sprintf("NM_MY_NFT_N?%s,no", value.TokenUid)),
			))
		}
	}
	keyboardUserNft.InlineKeyboard = append(keyboardUserNft.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_PROFILE"),
	))
	return keyboardUserNft
}

func GenKeyboardInlineForNftsToken(token []requestProject.Token, textBuy string, textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBuy, fmt.Sprintf("NM_MY_NFT_NSL?%s", token[0].TokenUid)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_MY_NFT"),
		),
	)
}

func GenKeyboardInlineForNftsTokenSell(textBackProfile string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBackProfile, "NM_MY_NFT"),
		),
	)
}

func GenKeyboardInlineForTokenSell(textBtnBack string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(textBtnBack, "NM_MY_NFT"),
		),
	)
}

func GenKeyboardInlineForAdminUserSellNft(uidPaymentEvent string, tokenUid string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Купить", fmt.Sprintf("NM_ADS?%s,%s", uidPaymentEvent, tokenUid)),
		),
	)
}

func GenKeyboardInlineForAdminUserWithDraw(eventWithDraw string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Одобрить", fmt.Sprintf("NM_WITH_DRAW_AA?%s", eventWithDraw)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("❌ Отказать", fmt.Sprintf("NM_WITH_DRAW_AR?%s", eventWithDraw)),
		),
	)
}
