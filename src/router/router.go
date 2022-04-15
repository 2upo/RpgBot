package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegrambot/utils"
)

const (
	Default Status = iota
	AddingNote
)


func InitRouter() Router{
	router := Router{}
	router.userSteps = make(map[Status]Handler)
	return router
}

func (this* Router)RegisterRoute(step Status, handler Handler){
	this.userSteps[]
}
