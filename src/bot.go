package main

import (
    "log"
    "sync"
    "telegrambot/utils"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// InitApp ...
func InitBot() (*tgbotapi.BotAPI) {
    config := utils.Config()

    // Instantiate TGBot instance
    bot, _ := tgbotapi.NewBotAPI(config.TgBotApiKey)
    bot.Debug = false

    log.Printf("Authorized on account %s", bot.Self.UserName)
    return bot
}

func RunBotListen(wg *sync.WaitGroup) {
    bot := InitBot()

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        log.Printf(update.Message.Text)
    }
}
