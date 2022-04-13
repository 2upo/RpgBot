package utils

type Status int

type Handler func(*tgbotapi.BotAPI, tgbotapi.Update) Status
