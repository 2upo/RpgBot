package main

import (
	"log"
	"sync"
	"telegrambot/utils"
)

var waitGroup sync.WaitGroup

func main() {
	// Instantiate database
	db := utils.Db()
	defer utils.CloseClient(10)
	log.Printf("Instantiated db: %v", db.Client())

	// Running Gin Server and Telegram bot API polling
	waitGroup.Add(2)
	RunServer(&waitGroup)
	RunBotListen(&waitGroup)
}
