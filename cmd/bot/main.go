package main

import (
	"log"
	"os"

	"github.com/Vladosya/nft-market-frontend/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	initConfig()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	if os.Getenv("IS_TESTING") == "true" {
		bot.Debug = true
	} else {
		bot.Debug = false
	}

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	for _, k := range []string{"TOKEN"} {
		if _, ok := os.LookupEnv(k); !ok {
			log.Fatal("set environment variable -> ", k)
		}
	}
}
