package main

import (
    "sync"
    "log"

	"telegrambot/utils"
    "telegrambot/router"
    
	"github.com/gin-gonic/gin"
)

// InitApp ...
func InitApp() (*gin.Engine) {
	// Gin Init
	app := gin.New()
	app.Use(gin.Recovery())

	// Router Init
	baseGroup := app.Group("/api")
	{
        // log.Printf("%v", baseGroup)
		router.InitRoutes(baseGroup)
	}

	return app
}

func RunServer(wg *sync.WaitGroup) {
    defer wg.Done()

    config := utils.Config()

	// Inicialize
	app := InitApp()

	log.Println("Starting application Webserver...")

	app.Run(config.ServerDsn)
}
