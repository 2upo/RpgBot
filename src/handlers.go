package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) Status {
	status := Default

	if update.Message == nil {
		return status
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "help":
		msg.Text = "Understandable!"

	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return status
}
