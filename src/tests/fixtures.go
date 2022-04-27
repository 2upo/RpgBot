package tests

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.9.0/mongo

// Fixture -- is a fabric function.
// Фабричная функция -- это (creational pattern) паттерн,
// который возвращает новый экземпляр какого-то объекта.

func ClearDb(collections []*mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	for _, collection := range collections {
		err := collection.Drop(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func newState(header string, stateCollection *mongo.Collection) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := stateCollection.InsertOne(ctx, bson.D{
		{"Header", header},
		{"Content", "Sample content"},
		{"CreatedAt", int(time.Now().Unix())},
		{"Answers", []bson.D{
			bson.D{
				{"NextState", "default"},
				{"Content", "sample content"},
			},
			bson.D{
				{"NextState", "default"},
				{"Content", "sample content"},
			},
		}},
	})
	return res, err
}

func NewUser(chatId string, userCollection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userCollection.InsertOne(ctx, bson.D{
		{"ChatId", chatId},
		{"CurrentState", primitive.NewObjectID()},
	})
	return res, err
}

func SetupStateCollection(stateCollection *mongo.Collection) []*mongo.InsertOneResult {

	state1, err := newState("state1", stateCollection)
	if err != nil {
		log.Fatal(err)
	}

	state2, err := newState("state2", stateCollection)
	if err != nil {
		log.Fatal(err)
	}

	state3, err := newState("state3", stateCollection)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(state3.InsertedID)

	// Example how to obtain state id:
	// id := res.InsertedID
	return []*mongo.InsertOneResult{state1, state2, state3}
}
