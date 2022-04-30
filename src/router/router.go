package router

import (
	"github.com/gin-gonic/gin"
	"telegrambot/state"
)

// InitRoutes ...
func InitRoutes(baseGroup *gin.RouterGroup) {
	baseGroup.GET("/healthcheck", healthcheck)

	// State instance
	stateController := state.InitController()
	stateController.InitRoutes(baseGroup)
}
