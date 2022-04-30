package main

import (
    "telegrambot/router"
	"telegrambot/tests"

    "net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func InitRoutes(app *gin.Engine) {
    baseGroup := app.Group(tests.TestAPIPath)
    {
        // State instance
        stateController := state.InitController()
        stateController.InitRoutes(baseGroup)

    }
}



func TestHealthcheck(t *testing.T) {
	app := tests.InitApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", tests.TestAPIPath + "/healthcheck", nil)
	app.ServeHTTP(w, req)

	response := &router.Healthcheck{
		Status: "healthy",
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response.String(), w.Body.String())
}

type MockStateService struct {}

func (s *MockStateService) GetAll() ([]State, error) {
    states := []State{
        State{
            
        }
    }
    return nil
}
    Insert(new_state *State) error
    GetById(id primitive.ObjectID) (*State, error)
    Update(updated_state *State) error
    DeleteById(id primitive.ObjectID) error
}
