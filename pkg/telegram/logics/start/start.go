package start

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func GetStart(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string) error {
	resCheckAuth, err := requestProject.CheckAuth(teleId)
	if err != nil {
		return err
	}
	fmt.Println("resCheckAuth -->", resCheckAuth)
	if resCheckAuth {
		fmt.Println("1")
		resGetUserLanguage, err := requestProject.GetUserLanguage(teleId)
		if err != nil {
			return err
		}
		fmt.Println("2")
		if len(resGetUserLanguage[0].Lang) > 0 {
			fmt.Println("3")
			resGetUserCurrency, err := requestProject.GetUserCurrency(teleId)
			if err != nil {
				return err
			}
			if len(resGetUserCurrency[0].Currency) > 0 {
				fmt.Println("4")
				resCheckIsTerms, err := requestProject.CheckIsTerms(teleId)
				if err != nil {
					return err
				}
				if resCheckIsTerms {
					fmt.Println("5")
					var textTwo string = ""
					if resGetUserLanguage[0].Lang == "ru" {
						textTwo = "Если у вас не появилось меню, то напишите /start"
					}
					if resGetUserLanguage[0].Lang == "en" {
						textTwo = "In case the menu did not appear, send /start or press it"
					}
					msg.Text = textTwo
					_, err = bot.Send(msg)
					if err != nil {
						return err
					}
					resCheckIsAdmin, err := requestProject.CheckIsAdmin(teleId)
					if err != nil {
						return err
					}
					if resCheckIsAdmin {
						if resGetUserLanguage[0].Lang == "ru" {
							msg.ReplyMarkup = keyboard.GenKeyboardHomeAdmin("NFT", "Личный кабинет", "Информация", "Поддержка")
							msg.Text = "Главное меню"
						}
						if resGetUserLanguage[0].Lang == "en" {
							msg.ReplyMarkup = keyboard.GenKeyboardHomeAdmin("NFT", "Profile", "About", "Support")
							msg.Text = "Main menu"
						}
					} else {
						if resGetUserLanguage[0].Lang == "ru" {
							msg.ReplyMarkup = keyboard.GenKeyboardHome("NFT", "Личный кабинет", "Информация", "Поддержка")
							msg.Text = "Главное меню"
						}
						if resGetUserLanguage[0].Lang == "en" {
							msg.ReplyMarkup = keyboard.GenKeyboardHome("NFT", "Profile", "About", "Support")
							msg.Text = "Main menu"
						}
					}
					_, err = bot.Send(msg)
					if err != nil {
						return err
					}
				} else {
					if resGetUserLanguage[0].Lang == "ru" {
						msg.ParseMode = "HTML"
						var text string = "Подтвердите, что вы не бот.\n\nНажмимая “Подтвердить“, Вы принимаете условия <a href='https://static.rarible.com/terms.pdf'>пользовательского соглашения</a>, <a href='https://static.rarible.com/privacy.pdf'>Условия конфиденциальности</a>."
						msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("✅ Подтвердить", true, resGetUserLanguage[0].Lang)
						msg.Text = text
						_, err := bot.Send(msg)
						if err != nil {
							return err
						}
					}
					if resGetUserLanguage[0].Lang == "en" {
						msg.ParseMode = "HTML"
						var text string = "Please, confirm you're not a robot.\n\nBy pressing “Accept“ you confirm that you've read and accept our <a href='https://static.rarible.com/terms.pdf'>Terms</a>, <a href='https://static.rarible.com/privacy.pdf'>Privacy</a>."
						msg.ReplyMarkup = keyboard.GenKeyboardInlineForAgreeTerms("✅ Accept", true, resGetUserLanguage[0].Lang)
						msg.Text = text
						_, err := bot.Send(msg)
						if err != nil {
							return err
						}
					}
				}
			} else {
				if resGetUserLanguage[0].Lang == "ru" {
					msg.Text = "Выберите свою валюту"
				}
				if resGetUserLanguage[0].Lang == "en" {
					msg.Text = "Choose your currency"
				}
				msg.ReplyMarkup = keyboard.DgCurrencyKeyboardInline
				_, err := bot.Send(msg)
				if err != nil {
					return err
				}
			}
		} else {
			var text string = "🏳️?"
			msg.ReplyMarkup = keyboard.DgLangKeyboardInline
			msg.Text = text
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	} else {
		resRegisterUser, err := requestProject.RegisterUser(teleId, userName)
		if err != nil {
			return err
		}
		if resRegisterUser {
			var text string = "🏳️?"
			msg.ReplyMarkup = keyboard.DgLangKeyboardInline
			msg.Text = text
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
