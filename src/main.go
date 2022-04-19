package main

import (
	"log"
	"telegrambot/utils"
  "telegrambot/router"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
  // Instantiate config
  config := utils.Config()

	// Instantiate database
	db := utils.Db()
	defer utils.CloseClient(10)
    log.Printf("Instantiated db: %v", db.Client)

	// Instantiate TGBot instance
	bot, _ := tgbotapi.NewBotAPI(config.TgBotApiKey)
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// TODO: logic here
	}
}
