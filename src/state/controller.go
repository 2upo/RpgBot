package stare

import (
    "telegrambot/utils"
)

type StateController struct {
    Service *StateService
}

// Constructor
func InitController() *StateController {
    var stateController StateController
    stateController.Service = InitStateService()

    return &stateController
}

func (controller *StateController)InitRoutes(baseGroup *gin.RouterGroup) {
    // Register routes
    stateGroup := baseGroup.Group("/state"){
        stateGroup.POST("/", controller.insertState)
    }
}

func (controller *StateController)insertState(ctx *gin.Context) {
    var state State
    if err := ctx.ShouldBindJSON(&state); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := controller.Service.Insert(&state); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    } else {
        ctx.JSON(http.StatusOK, gin.H{
    		"Success": state,
    	})
    }
}
