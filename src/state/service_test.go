package state

import (
	"context"
	"telegrambot/tests"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetAll(t *testing.T) {
	ass := assert.New(t)

	tests.SetupStateCollection(collection) // Set up

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

	tests.ClearDb([]*mongo.Collection{collection}) // Tear down
}

// Test GetAll against empty Collection
func TestGetAllEmpty(t *testing.T) {
	ass := assert.New(t)

	states, err := GetAll()
	if err != nil {
		ass.Nil(err)
	}

	ass.Equal(len(states), 0)

	tests.ClearDb([]*mongo.Collection{collection}) // Tear down
}

func TestGetById(t *testing.T) {
	ass := assert.New(t)

	states := tests.SetupStateCollection(collection) // Set up

	state, err := GetById(states[0].InsertedID.(primitive.ObjectID))

	ass.Nil(err)
	ass.Equal(state.Header, "state1")

	tests.ClearDb([]*mongo.Collection{collection}) // Tear down
}

// Test GetById against unexisting id
func TestGetByEmptyId(t *testing.T) {
	ass := assert.New(t)

	_, err := GetById(primitive.NewObjectID())

	ass.Equal(err, mongo.ErrNoDocuments)

	tests.ClearDb([]*mongo.Collection{collection}) // Tear down
}

func TestInsert(t *testing.T) {
	ass := assert.New(t)

	new_state := State{
		Header:  "state1",
		Content: "sample content",
	}

	err := Insert(&new_state)
	ass.Nil(err)

	state, err := mockGetById(new_state.ID)

	ass.Nil(err)
	ass.Equal(state.Header, new_state.Header)
	ass.Equal(state.Content, new_state.Content)

	tests.ClearDb([]*mongo.Collection{collection})
}

func TestUpdate(t *testing.T) {
	ass := assert.New(t)

	states := tests.SetupStateCollection(collection)

	state, err := mockGetById(states[0].InsertedID.(primitive.ObjectID))
	ass.Nil(err)

	state.Content = "new content"

	err = Update(state)
	ass.Nil(err)

	updatedState, err := mockGetById(states[0].InsertedID.(primitive.ObjectID))
	ass.Nil(err)
	ass.Equal(updatedState.Content, state.Content)

	tests.ClearDb([]*mongo.Collection{collection})
}

func TestDelete(t *testing.T) {
	ass := assert.New(t)
	states := tests.SetupStateCollection(collection)

	err := DeleteById(states[0].InsertedID.(primitive.ObjectID))
	ass.Nil(err)

	_, err = mockGetById(states[0].InsertedID.(primitive.ObjectID))
	ass.Equal(err, mongo.ErrNoDocuments)
	state, err := mockGetById(states[1].InsertedID.(primitive.ObjectID))
	ass.Nil(err)
	ass.NotNil(state)
	tests.ClearDb([]*mongo.Collection{collection})
}

func mockGetById(id primitive.ObjectID) (*State, error) {
	var state State

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&state)

	return &state, err
}
