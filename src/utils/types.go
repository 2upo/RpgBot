package utils

import (
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Current user status type: corresponds to constants
type Status int

// Delegate describes uodate message handler
type Handler func(*tgbotapi.BotAPI, tgbotapi.Update) Status
