package telegram

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/helperFunc"
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
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/addBalance"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/addMamMinim"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/addMammoth"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/blockUser"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/changeBalance"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/changeMamMinimal"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/changeMamPremium"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/changeMamVerification"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/findMammoth"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/mammothProfile"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/mammothsShow"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/logics/workerPanel/messageMammoth"
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
		resCheckIsAdmin, err := requestProject.CheckIsAdmin(message.From.ID)
		if err != nil {
			return err
		}
		if resCheckIsAdmin {
			err = sqlite.TurnOffListeners(b.SqliteDb, message.From.ID)
			if err != nil {
				return err
			}
			msg.ParseMode = "Markdown"
			err := start.GetStart(b.Bot, msg, message.From.ID, message.From.UserName, message.From.ID)
			if err != nil {
				return err
			}
		} else {
			if len(needParams) > 1 {
				i, err := strconv.ParseInt(needParams[1], 10, 64)
				if err != nil {
					return err
				}
				resCheckIsAdmin, err = requestProject.CheckIsAdmin(i)
				if err != nil {
					return err
				}
				if resCheckIsAdmin {
					msg.ParseMode = "Markdown"
					err := start.GetStart(b.Bot, msg, message.From.ID, message.From.UserName, i)
					if err != nil {
						return err
					}
				}
			} else {
				resCheckAuth, err := requestProject.CheckAuth(message.From.ID)
				if err != nil {
					return err
				}
				if resCheckAuth {
					resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.From.ID)
					if err != nil {
						return err
					}
					if !resCheckIsBlockUser {
						err := start.GetStart(b.Bot, msg, message.From.ID, message.From.UserName, message.From.ID)
						if err != nil {
							return err
						}
					}
				} else {
					msg.ParseMode = "Markdown"
					msg.Text = "⚠️ Вход доступен только по реферальной ссылке"
					_, err := b.Bot.Send(msg)
					if err != nil {
						return err
					}
				}
			}
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
				err = mammothProfile.MammothProfile(b.Bot, b.SqliteDb, msg, message.From.ID, message.From.UserName, resGetUserLang, i)
				if err != nil {
					return err
				}
			}
		}
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
		err = workerPanel.WorkerPanel(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
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
		err = mammothProfile.MammothProfile(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_SRC":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = findMammoth.FindMammoth(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_CMIN":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = changeMamMinimal.ChangeMamMinimal(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_MAM":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = myMammoths.MyMammoths(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_SF":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.Atoi(needParams[1])
		if err != nil {
			return err
		}
		err = mammothsShow.MammothsShow(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_AD":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		err = addMammoth.AddMammoth(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang)
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
	case "NM_WORKPANEL_MAM_VERIF":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = changeMamVerification.ChangeMamVerification(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_ADB":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = addBalance.AddBalance(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_MNU":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = addMamMinim.AddMamMinim(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_CHB":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = changeBalance.ChangeBalance(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_MSM":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = messageMammoth.MessageMammoth(b.Bot, b.SqliteDb, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
		if err != nil {
			return err
		}
	case "NM_WORKPANEL_BUS":
		resGetUserLang, err := sqlite.GetUserLang(b.SqliteDb, callbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		i, err := strconv.ParseInt(needParams[1], 10, 64)
		if err != nil {
			return err
		}
		err = blockUser.BlockUser(b.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName, resGetUserLang, i)
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

	resCheckIsListenerFindM, err := sqlite.CheckIsListenerAddM(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	if resCheckIsListenerFindM {
		var params = strings.Split(message.Text, " ")
		if len(params) == 2 {
			_, err := strconv.Atoi(params[0])
			if err != nil {
				msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: _1234567890 Джони_"
				msg.ReplyMarkup = keyboard.DgAddMammothKeyboardInline
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
			}
			i, err := strconv.ParseInt(params[0], 10, 64)
			if err != nil {
				return err
			}
			_, err = requestProject.CreateReferral(i, params[1], message.Chat.ID)
			if err != nil {
				return err
			}
			err = sqlite.TurnOffListeners(b.SqliteDb, message.Chat.ID)
			if err != nil {
				return err
			}
			msg.Text = fmt.Sprintf("✅ Мамонт с Id %d был успешно добавлен", i)
			msg.ReplyMarkup = keyboard.DgAddMammothSuccessKeyboardInline
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: _1234567890 Джони_"
			msg.ReplyMarkup = keyboard.DgAddMammothKeyboardInline
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}
	resCheckIsListenerFindM, err = sqlite.CheckIsListenerFindM(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}

	if resCheckIsListenerFindM {
		_, err := strconv.Atoi(message.Text)
		if err != nil {
			msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: _1234567890_"
			msg.ReplyMarkup = keyboard.DgFindMammothKeyboardInline
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			i, err := strconv.ParseInt(message.Text, 10, 64)
			if err != nil {
				return err
			}
			resGetUserReferral, err := requestProject.GetUserReferral(message.Chat.ID, i)
			if err != nil {
				return err
			}
			if len(resGetUserReferral) > 0 {
				msg.Text = fmt.Sprintf("✅ Мамонт с Id %s был успешно найден", message.Text)
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}

				msg.ParseMode = "HTML"
				resModifyDate, err := helperFunc.ModifyDate(resGetUserReferral[0].Created)
				if err != nil {
					return err
				}
				msg.Text = fmt.Sprintf("🦣 %s\n📅 %s\n📂 <a href='%d'>/u%d</a>\n\n", resGetUserReferral[0].TeleName, resModifyDate, resGetUserReferral[0].TeleId, resGetUserReferral[0].TeleId)
				msg.ReplyMarkup = keyboard.DgFindMammothSuccessKeyboardInline
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
				err = sqlite.TurnOffListeners(b.SqliteDb, message.Chat.ID)
				if err != nil {
					return err
				}
			} else {
				msg.Text = "⛔️ Пользователь с данным идентификатором не был найден.\n\nФормат: _1234567890_"
				msg.ReplyMarkup = keyboard.DgFindMammothKeyboardInline
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
			}
		}
	}

	resCheckIsListenerWatchingChangeMinLink, err := sqlite.CheckIsListenerWatchingChangeMinLink(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	if resCheckIsListenerWatchingChangeMinLink {
		i, err := strconv.ParseFloat(message.Text, 64)
		if err != nil {
			msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: только числовые данные"
			msg.ReplyMarkup = keyboard.DgChangeMamMinimalKeyboardInline
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resAdminUpdateMinimPrice, err := requestProject.AdminUpdateMinimPrice(message.Chat.ID, i)
			if err != nil {
				return err
			}
			if resAdminUpdateMinimPrice {
				msg.Text = "✅ Успешное изменение минималки"
				msg.ReplyMarkup = keyboard.DgChangeMamMinimalKeyboardInline
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
				err = sqlite.TurnOffListeners(b.SqliteDb, message.Chat.ID)
				if err != nil {
					return err
				}
			}
		}
	}
	resCheckIsListenerWatchingAddBalance, teleIdUser, err := sqlite.CheckIsListenerWatchingAddBalance(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	if resCheckIsListenerWatchingAddBalance {
		i, err := strconv.ParseFloat(message.Text, 64)
		if err != nil {
			msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: только числовые данные"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUser)
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resAdminAddBalance, err := requestProject.AdminAddBalance(teleIdUser, i)
			if err != nil {
				return err
			}
			if resAdminAddBalance {
				msg.ChatID = teleIdUser
				if resGetUserLang == "ru" {
					msg.Text = fmt.Sprintf("На ваш аккаунт добавлено: *%.2f$*", i)
				}
				if resGetUserLang == "en" {
					msg.Text = fmt.Sprintf("Added to your account: *%.2f$*", i)
				}
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
				msg.ChatID = message.Chat.ID
				msg.Text = "✅ Успешное добавление баланса мамонту"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUser)
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
				err = sqlite.TurnOffListeners(b.SqliteDb, message.Chat.ID)
				if err != nil {
					return err
				}
			}
		}
	}

	resCheckIsListenerWatchingAddUser, teleIdUserTwo, err := sqlite.CheckIsListenerWatchingAddMinUser(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	if resCheckIsListenerWatchingAddUser {
		i, err := strconv.ParseFloat(message.Text, 64)
		if err != nil {
			msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: только числовые данные"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUserTwo)
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resAdminChangeMinUser, err := requestProject.AdminChangeMinUser(teleIdUserTwo, i)
			if err != nil {
				return err
			}
			if resAdminChangeMinUser {
				msg.ChatID = message.Chat.ID
				msg.Text = "✅ Успешное добавление минималки для мамонта"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUserTwo)
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
				err = sqlite.TurnOffListeners(b.SqliteDb, message.Chat.ID)
				if err != nil {
					return err
				}
			}
		}
	}

	resCheckIsListenerWatchingChangeBalance, teleIdUserThree, err := sqlite.CheckIsListenerWatchingChangeBalance(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	if resCheckIsListenerWatchingChangeBalance {
		i, err := strconv.ParseFloat(message.Text, 64)
		if err != nil {
			msg.Text = "⛔️ Невалидный формат введённых данных!\n\nФормат: только числовые данные"
			msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUserThree)
			_, err = b.Bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			resAdminChangeMinUser, err := requestProject.AdminChangeBalance(teleIdUserThree, i)
			if err != nil {
				return err
			}
			if resAdminChangeMinUser {
				msg.ChatID = message.Chat.ID
				msg.Text = "✅ Успешное изменение баланса для мамонта"
				msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUserThree)
				_, err = b.Bot.Send(msg)
				if err != nil {
					return err
				}
				err = sqlite.TurnOffListeners(b.SqliteDb, message.Chat.ID)
				if err != nil {
					return err
				}
			}
		}
	}

	resCheckIsListenerWatchingMessageUser, teleIdUserFour, err := sqlite.CheckIsListenerWatchingMessageUser(b.SqliteDb, message.Chat.ID)
	if err != nil {
		return err
	}
	if resCheckIsListenerWatchingMessageUser {
		msg.ChatID = teleIdUserFour
		msg.Text = message.Text
		_, err = b.Bot.Send(msg)
		if err != nil {
			return err
		}

		msg.ChatID = message.Chat.ID
		msg.Text = "✅ Успешная отправка сообщению мамонту"
		msg.ReplyMarkup = keyboard.GenKeyboardInlineForAddBalance(teleIdUserFour)
		_, err = b.Bot.Send(msg)
		if err != nil {
			return err
		}
	}

	switch message.Text {
	case "NFT 🎆":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := nft.Nft(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "Profile 📁":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := profile.Profile(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "About ℹ️":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := about.About(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "👨‍💻 Support":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := support.Support(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "Личный кабинет 📁":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := profile.Profile(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "Информация ℹ️":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := about.About(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "👨‍💻 Поддержка":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := support.Support(b.Bot, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	case "🛠 Панель воркера":
		resCheckIsBlockUser, err := requestProject.CheckIsBlockUser(message.Chat.ID)
		if err != nil {
			return err
		}
		if !resCheckIsBlockUser {
			err := workerPanel.WorkerPanel(b.Bot, b.SqliteDb, msg, message.Chat.ID, message.Chat.UserName, resGetUserLang)
			if err != nil {
				return err
			}
		} else {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
			_, err := b.Bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
