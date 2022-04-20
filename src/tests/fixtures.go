package tests

import (
    "context"
    "time"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"

    "telegrambot/utils"
)
// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.9.0/mongo

// Fixture -- is a fabric function.
// Фабричная функция -- это (creational pattern) паттерн,
// который возвращает новый экземпляр какого-то объекта.

var db = utils.Db()

func ClearDb() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    err := db.Collection("state").Drop(ctx)
    if err != nil {
        log.Fatal(err)
    }
}

func newState(header string) (*mongo.InsertOneResult, error) {
    collection := db.Collection("state")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := collection.InsertOne(ctx, bson.D{
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

func SetupStateCollection() []*mongo.InsertOneResult {

    state1, err := newState("state1")
    if err != nil {
        log.Fatal(err)
    }

    state2, err := newState("state2")
    if err != nil {
        log.Fatal(err)
    }

    state3, err := newState("state3")
    if err != nil {
        log.Fatal(err)
    }

    log.Println(state3.InsertedID)

    // Example how to obtain state id:
    // id := res.InsertedID
    return []*mongo.InsertOneResult{state1, state2, state3}
}
