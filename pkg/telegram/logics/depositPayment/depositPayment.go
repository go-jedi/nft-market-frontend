package depositPayment

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
)

func DepositPayment(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string, userPaymentChoose string, userPriceWrite float64) error {
	if len(languageUser) > 0 {
		if len(userPaymentChoose) > 0 {
			resGetAdminByUser, err := requestProject.GetAdminByUser(teleId)
			if err != nil {
				return err
			}
			if len(resGetAdminByUser) > 0 {
				resCreateDepot, err := requestProject.CreateDepot(teleId, userName, resGetAdminByUser[0].TeleId, resGetAdminByUser[0].TeleName, userPriceWrite)
				if err != nil {
					return err
				}
				err = requestProject.SendTelegramNewDepot()
				if err != nil {
					return err
				}
				if resCreateDepot {
					resGetAllPayments, err := requestProject.GetAllPayments()
					if err != nil {
						return err
					}
					if len(resGetAllPayments) > 0 {
						resGetAllExchangeRates, err := requestProject.GetAllExchangeRates()
						if err != nil {
							return err
						}
						if len(resGetAllExchangeRates) > 0 {
							resGetUserMinPrice, err := requestProject.GetUserMinPrice(teleId)
							if err != nil {
								return err
							}
							if len(resGetUserMinPrice) > 0 {
								if languageUser == "ru" {
									if userPaymentChoose == "btc" {
										var text string = fmt.Sprintf("Оплата в BTC\n\nЧтобы пополнить свой BTC из внешнего кошелька, используйте многоразовый адрес ниже.\n\n💱 Адрес BTC: *%s*\nСумма пополнения: *%.4f btc*\n\nПосле внесения средств отправьте скриншот перевода в службу технической поддержки, и средства будут зачислены на ваш счет.\n\n⚠️ Уважаемый пользователь, обратите внимание, что все платежи менее *%.2f $* не будут зачислены на сервис, компенсация за данные операции также не предусмотрена.",
											resGetAllPayments[0].Value,
											resGetAllExchangeRates[0].Price*userPriceWrite,
											resGetUserMinPrice[0].MinimPrice,
										)
										msg.Text = text
										msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositPayment("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
										_, err := bot.Send(msg)
										if err != nil {
											return err
										}
									}
									if userPaymentChoose == "eth" {
										var text string = fmt.Sprintf("Оплата в ETH\n\nЧтобы пополнить свой ETH из внешнего кошелька, используйте многоразовый адрес ниже.\n\n💱 Адрес ETH:\n*%s*\nСумма пополнения: *%.4f eth*\n\nПосле внесения средств отправьте скриншот перевода в службу технической поддержки, и средства будут зачислены на ваш счет.\n\n⚠️ Уважаемый пользователь, обратите внимание, что все платежи менее *%.2f $* не будут зачислены на сервис, компенсация за данные операции также не предусмотрена.",
											resGetAllPayments[1].Value,
											resGetAllExchangeRates[1].Price*userPriceWrite,
											resGetUserMinPrice[0].MinimPrice,
										)
										msg.Text = text
										msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositPayment("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
										_, err := bot.Send(msg)
										if err != nil {
											return err
										}
									}
									if userPaymentChoose == "usdt" {
										var text string = fmt.Sprintf("Оплата в USDT\n\nЧтобы пополнить свой USDT из внешнего кошелька, используйте многоразовый адрес ниже.\n\n💱 Адрес USDT:\n*%s*\nСумма пополнения: *%.0f $*\n\nПосле внесения средств отправьте скриншот перевода в службу технической поддержки, и средства будут зачислены на ваш счет.\n\n⚠️ Уважаемый пользователь, обратите внимание, что все платежи менее *%.2f $* не будут зачислены на сервис, компенсация за данные операции также не предусмотрена.",
											resGetAllPayments[2].Value,
											resGetAllExchangeRates[2].Price*userPriceWrite,
											resGetUserMinPrice[0].MinimPrice,
										)
										msg.Text = text
										msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositPayment("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
										_, err := bot.Send(msg)
										if err != nil {
											return err
										}
									}
								}

								if languageUser == "en" {
									if userPaymentChoose == "btc" {
										var text string = fmt.Sprintf("Payment in BTC\n\nTo top up your BTC from an external wallet, use the reusable address below.\n\n💱 BTC Address: *%s*\n\nAfter depositing funds, send a screenshot of the transfer to technical support and the funds will be credited to your account.\n\n⚠️ Dear user, please note that all entries less than *%.2f $* will not be credited to the service, compensation for these operations is also not provided.",
											resGetAllPayments[0].Value,
											resGetUserMinPrice[0].MinimPrice,
										)
										msg.Text = text
										msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositPayment("👨‍💻 Support", "🔙 Back to profile")
										_, err := bot.Send(msg)
										if err != nil {
											return err
										}
									}
									if userPaymentChoose == "eth" {
										var text string = fmt.Sprintf("Payment in ETH\n\nTo top up your ETH from an external wallet, use the reusable address below.\n\n💱 ETH Address:\n*%s*\n\nAfter depositing funds, send a screenshot of the transfer to technical support and the funds will be credited to your account.\n\n⚠️ Dear user, please note that all entries less than *%.2f ₽* will not be credited to the service, compensation for these operations is also not provided.",
											resGetAllPayments[1].Value,
											resGetUserMinPrice[0].MinimPrice,
										)
										msg.Text = text
										msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositPayment("👨‍💻 Support", "🔙 Back to profile")
										_, err := bot.Send(msg)
										if err != nil {
											return err
										}
									}
									if userPaymentChoose == "usdt" {
										var text string = fmt.Sprintf("Payment in USDT\n\nTo top up your USDT from an external wallet, use the reusable address below.\n\n💱 USDT Address:\n*%s*\n\nAfter depositing funds, send a screenshot of the transfer to technical support and the funds will be credited to your account.\n\n⚠️ Dear user, please note that all entries less than *%.2f ₽* will not be credited to the service, compensation for these operations is also not provided.",
											resGetAllPayments[2].Value,
											resGetUserMinPrice[0].MinimPrice,
										)
										msg.Text = text
										msg.ReplyMarkup = keyboard.GenKeyboardInlineForDepositPayment("👨‍💻 Support", "🔙 Back to profile")
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
			}
		}
	}

	return nil
}
