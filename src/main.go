package main

import (
	"context"
	"log"
	"telegrambot/utils"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Instantiate database
	db := utils.Db()

	// Instantiate TGBot instance
	bot, err := tgbotapi.NewBotAPI("")

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	// TODO: suda cod pishi
	var userStatus UserStatuses
	userStatus.InitMap()

	for update := range updates {
		userId := update.Message.Chat.ID
		handlerToProcess := userStatus.GetCurrentUserHandler(userId)
		newStatus := handlerToProcess(bot, update)
		userStatus.SetUserStatus(userId, newStatus)
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println(newStatus)
	}
}
