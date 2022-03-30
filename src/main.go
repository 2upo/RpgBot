package main

import (
	"context"
	"fmt"
	"log"
	"telegrambot/utils"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	db := utils.Db()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	err := db.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	cancel()

	//Telegram
	bot, err := tgbotapi.NewBotAPI("5201221376:AAE6eVSgAIsRbrta-9ujkY347xq33B15-2c")
	CheckError(err)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
