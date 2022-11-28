package nftTokenBuy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func NftTokenBuy(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userTokenChoose string) error {
	if len(languageUser) > 0 {
		if len(userTokenChoose) > 0 {
			resGetToken, err := requestProject.GetToken(userTokenChoose)
			if err != nil {
				return err
			}
			resCheckUserToken, err := requestProject.CheckUserToken(teleId, userTokenChoose)
			if err != nil {
				return err
			}
			if resCheckUserToken {
				if languageUser == "ru" {
					msg.Text = "❕ У вас уже есть этот NFT."
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuyHaveToken(resGetToken, "🔙 Назад")
				}
				if languageUser == "en" {
					msg.Text = "❕ You already have this NFT."
					msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuyHaveToken(resGetToken, "🔙 Back")
				}
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}

			if !resCheckUserToken {
				if len(resGetToken) > 0 {
					resGetUserBalance, err := requestProject.GetUserBalance(teleId)
					if err != nil {
						return err
					}
					if len(resGetUserBalance) > 0 {
						if resGetUserBalance[0].Balance >= resGetToken[0].Price {
							resBuyUserToken, err := requestProject.BuyUserToken(teleId, resGetToken[0].TokenUid, resGetToken[0].Price)
							if err != nil {
								return err
							}
							if resBuyUserToken {
								if languageUser == "ru" {
									msg.Text = "✅ Токен был успешно куплен."
									msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuy("🔙 Назад", userTokenChoose)
								}
								if languageUser == "en" {
									msg.Text = "✅ The token has been successfully purchased."
									msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuy("🔙 Back", userTokenChoose)
								}
								_, err := bot.Send(msg)
								if err != nil {
									return err
								}
							}
						}

						if resGetUserBalance[0].Balance < resGetToken[0].Price {
							if languageUser == "ru" {
								msg.Text = "❌ На вашем счете недостаточно средств для покупки токена."
								msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuy("🔙 Назад", userTokenChoose)
							}
							if languageUser == "en" {
								msg.Text = "❌ You don't have enough funds to buy this token."
								msg.ReplyMarkup = keyboard.GenKeyboardInlineForNftTokenBuy("🔙 Back", userTokenChoose)
							}
							_, err := bot.Send(msg)
							if err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}
