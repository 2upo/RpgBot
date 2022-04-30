package state

import (
	"context"
	"telegrambot/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StateService struct {
	Collection *mongo.Collection
}

// Constructor
func InitStateService() *StateService {
	var stateService StateService
	stateService.Collection = utils.Db().Collection("state")

	return &stateService
}

// Get all states from database
func (service *StateService) GetAll() ([]State, error) {
	var states []State

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// defining select options
	options := options.Find()
	options.SetSort(bson.D{{"createdat", -1}})

	// find all
	// TODO: add pagination
	// options.SetLimit(10)
	// options.SetSkip(10)

	cursor, err := service.Collection.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result State

		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		states = append(states, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return states, nil
}

func (service *StateService) Insert(new_state *State) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Define created at timestamp:
	new_state.CreatedAt = int(time.Now().Unix())

	res, err := service.Collection.InsertOne(ctx, new_state)
	if err == nil {
		new_state.ID = res.InsertedID.(primitive.ObjectID)
	}
	return err
}

func (service *StateService) GetById(id primitive.ObjectID) (*State, error) {
	var state State

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := service.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&state)

	return &state, err
}

func (service *StateService) Update(updated_state *State) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := service.Collection.ReplaceOne(ctx, bson.M{"_id": updated_state.ID}, updated_state)

	return err
}

func (service *StateService) DeleteById(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := service.Collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}
