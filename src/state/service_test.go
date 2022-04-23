package state

import (
	"testing"
	"telegrambot/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAll(t *testing.T) {
	ass := assert.New(t)

	tests.SetupStateCollection() // Set up

	states, err := GetAll()
	if err != nil {
		ass.Nil(err)
	}

	ass.Equal(len(states), 3)

	// Checking state content:
	ass.Equal(states[0].Content, "Sample content")
	ass.NotNil(states[0].Answers[0].NextState)
	ass.Equal(states[0].Answers[0].Content, "sample content")
	ass.NotNil(states[0].ID)

	tests.ClearDb() // Tear down
}

func TestGetById(t *testing.T) {
	ass := assert.New(t)

	states := tests.SetupStateCollection() // Set up

	state, err := GetById(states[0].InsertedID.(primitive.ObjectID))

	ass.Nil(err)
	ass.Equal(state.Header, "state1")

	tests.ClearDb() // Tear down
}
