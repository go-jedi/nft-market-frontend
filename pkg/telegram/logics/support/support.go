package support

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram/keyboard"
)

func Support(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, teleId int64, userName string, languageUser string) error {
	var isTesting string = os.Getenv("IS_TESTING")
	var needPath string = ""
	if isTesting == "true" {
		needPath = "/home/dale/job/work/my-project/nft-market/frontend/img"
	} else {
		needPath = "/home/nft-market-bot/frontend/nft-market-frontend/img"
	}
	photo := tgbotapi.NewPhoto(teleId, tgbotapi.FilePath(fmt.Sprintf("%s%s", needPath, "/img-need/3.jpg")))
	photo.ParseMode = "Markdown"
	if languageUser == "ru" {
		photo.Caption = "*Правила обращения в Техническую Поддержку:*\n\n🔹1. *Представьтесь, изложите проблему своими словами* - мы постараемся Вам помочь.\n\n🔹2. *Напишите свой ID* - нам это нужно, чтобы увидеть ваш профиль, и узнать актуальность вашей проблемы.\n\n🔹3. *Будьте вежливы*, наши консультанты не роботы, мы постараемся помочь Вам, и сделать все возможное, чтобы сберечь ваше время и обеспечить максимальную оперативность в работе.\n\n_Напишите нам, ответ Поддержки, не заставит вас долго ждать!_"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForSupportMenu("👨‍💻 Поддержка", "🔙 Вернуться в ЛК")
	}
	if languageUser == "en" {
		photo.Caption = "*Rules for contacting Technical Support:*\n\n🔹1. *Introduce yourself*, state the problem in your own words - we will try to help you.\n\n🔹2. *Send us your ID* - we need this to see your profile and see if your issue is up to date.\n\n🔹3. *Be polite*, our consultants are not robots, we will try to help you and do our best to save your time and ensure maximum efficiency in work.\n\n_Message us, our Support team will not keep you waiting long!_"
		photo.ReplyMarkup = keyboard.GenKeyboardInlineForSupportMenu("👨‍💻 Support", "🔙 Back to profile")
	}
	_, err := bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}
