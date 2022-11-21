package telegram

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/about"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/agreeTerms"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/changeLang"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/chooseCurrency"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/deposit"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/depositPayment"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/homeAfterReg"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/myNfts"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/nft"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/nftCollection"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/nftToken"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/nftTokenBuy"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/profile"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/start"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/support"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/verification"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/withDraw"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/withDrawPayment"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/changeMamPremium"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/mammothProfile"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/myMammoths"
	requestProject "github.com/rob-bender/nft-market-frontend/pkg/telegram/request"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/sqlite"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.From.ID, "Я не знаю такой команды")
	var needParams = strings.Split(message.Text, " ")
	var needParamsReferr = strings.Split(message.Text, "/u")

	switch message.Command() {
	case CommandStart:
		if len(needParams) > 1 {
			i, err := strconv.ParseInt(needParams[1], 10, 64)
			if err != nil {
				return err
			}
			_, err = requestProject.CreateReferral(message.From.ID, message.From.UserName, i)
			if err != nil {
				return err
			}
		}
		msg.ParseMode = "Markdown"
		fmt.Println("start111111111111111111")
		err := start.GetStart(b.Bot, msg, message.From.ID, message.From.UserName)
		if err != nil {
			return err
		}
	case CommandHelp:
		msg.Text = "Ты ввёл команду /help"
		_, err := b.Bot.Send(msg)
		if err != nil {
			return err
		}
	default:
		if len(needParamsReferr) == 2 {
			resCheckIsAdmin, err := requestProject.CheckIsAdmin(message.From.ID)
			if err != nil {
				return err
			}
			if resCheckIsAdmin {
				resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, message.From.ID)
				if err != nil {
					return err
				}
				i, err := strconv.ParseInt(needParamsReferr[1], 10, 64)
				if err != nil {
					return err
				}
				err = mammothProfile.MammothProfile(b.Bot, msg, message.From.ID, message.From.UserName, resGetUserLang, i)
				if err != nil {
					return err
				}
			}
		}
		// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
		// _, err := b.Bot.Send(msg)
		// if err != nil {
		// 	return err
		// }
	}

	return nil
}

func (b *Bot) callbackQuery(callbackQuery tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "callbackQuery")
	msg.ParseMode = "Markdown"

	var needParams = strings.Split(callbackQuery.Data, "?")
	fmt.Println("callbackQuery.Data -->", callbackQuery.Data)
	if len(needParams) > 1 {
		callbackQuery.Data = needParams[0]
	}

	switch callbackQuery.Data {
	case "NM_LANGUAGE":
		err := chooseCurrency.ChooseCurrency(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, needParams[1])
		if err != nil {
			return err
		}
	case "NM_CURRENCY":
		err := agreeTerms.AgreeTerms(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, needParams[1])
		if err != nil {
			return err
		}
	case "DG_AREG_HOME":
		var params = strings.Split(needParams[1], ",")
		if len(params) == 2 {
			i, err := strconv.ParseBool(params[0])
			if err != nil {
				return err
			}
			err = homeAfterReg.HomeAfterReg(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, i, params[1])
			if err != nil {
				return err
			}
			resCheckIsAdmin, err := requestProject.CheckIsAdmin(callbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			if resCheckIsAdmin {
				if params[1] == "ru" {
					msg.ReplyMarkup = keyboard.GenKeyboardHomeAdmin("NFT", "Личный кабинет", "Информация", "Поддержка")
					msg.Text = "Главное меню"
				}
				if params[1] == "en" {
					msg.ReplyMarkup = keyboard.GenKeyboardHomeAdmin("NFT", "Profile", "About", "Support")
					msg.Text = "Main menu"
				}
			} else {
				if params[1] == "ru" {
					msg.ReplyMarkup = keyboard.GenKeyboardHome("NFT", "Личный кабинет", "Информация", "Поддержка")
					msg.Text = "Главное меню"
				}
				if params[1] == "en" {
					msg.ReplyMarkup = keyboard.GenKeyboardHome("NFT", "Profile", "About", "Support")
					msg.Text = "Main menu"
				}
			}
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "NM_NFT":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = nft.Nft(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_PROFILE":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = profile.Profile(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_TOP_UP":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = deposit.Deposit(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_DEPOSIT_PAYMT":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = depositPayment.DepositPayment(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, needParams[1])
		if err != nil {
			return err
		}
	case "NM_MY_NFT":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = myNfts.MyNfts(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_VERIF":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = verification.Verification(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_CHNG_LANG":
		resCheckIsAdmin, err := requestProject.CheckIsAdmin(callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		if resCheckIsAdmin {
			if needParams[1] == "ru" {
				msg.ReplyMarkup = keyboard.GenKeyboardHomeAdmin("NFT", "Личный кабинет", "Информация", "Поддержка")
				msg.Text = "Главное меню"
			}
			if needParams[1] == "en" {
				msg.ReplyMarkup = keyboard.GenKeyboardHomeAdmin("NFT", "Profile", "About", "Support")
				msg.Text = "Main menu"
			}
		} else {
			if needParams[1] == "ru" {
				msg.ReplyMarkup = keyboard.GenKeyboardHome("NFT", "Личный кабинет", "Информация", "Поддержка")
				msg.Text = "Главное меню"
			}
			if needParams[1] == "en" {
				msg.ReplyMarkup = keyboard.GenKeyboardHome("NFT", "Profile", "About", "Support")
				msg.Text = "Main menu"
			}
		}
		_, err = b.Bot.Send(msg)
		if err != nil {
			return err
		}
		err = changeLang.ChangeLang(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, needParams[1])
		if err != nil {
			return err
		}
	case "NM_WITH_DRAW":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = withDraw.WithDraw(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_WITH_DRAW_PAYMT":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = withDrawPayment.WithDrawPayment(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, needParams[1])
		if err != nil {
			return err
		}
	case "NM_NFT_COLL":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = nftCollection.NftCollection(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, needParams[1])
		if err != nil {
			return err
		}
	case "NM_NFT_COLL_T":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = nftToken.NftToken(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, needParams[1])
		if err != nil {
			return err
		}
	case "NM_NFT_COLL_TB":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = nftTokenBuy.NftTokenBuy(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, needParams[1])
		if err != nil {
			return err
		}
	case "NM_WORKPANEL":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = workerPanel.WorkerPanel(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_MAM_US":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = mammothProfile.MammothProfile(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_MAM":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = myMammoths.MyMammoths(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_MAM_PREM":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = changeMamPremium.ChangeMamPremium(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ParseMode = "Markdown"
	resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	switch message.Text {
	case "NFT 🎆":
		err := nft.Nft(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "Profile 📁":
		err := profile.Profile(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "About ℹ️":
		err := about.About(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "👨‍💻 Support":
		err := support.Support(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "Личный кабинет 📁":
		err := profile.Profile(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "Информация ℹ️":
		err := about.About(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "👨‍💻 Поддержка":
		err := support.Support(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "🛠 Панель воркера":
		err := workerPanel.WorkerPanel(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	}

	return nil
}
