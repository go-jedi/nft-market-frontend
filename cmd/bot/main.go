package main

import (
	"database/sql"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rob-bender/nft-market-frontend/pkg/telegram"
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

	db, err := sql.Open("sqlite3", "my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	telegramBot := telegram.NewBot(bot, db)
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
