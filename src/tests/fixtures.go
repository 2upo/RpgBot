package tests

import (
    "context"
    "time"
    "log"

    "go.mongodb.org/mongo-driver/bson"

    "telegrambot/utils"
)
// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.9.0/mongo

// Fixture -- is a fabric function.
// Фабричная функция -- это (creational pattern) паттерн,
// который возвращает новый экземпляр какого-то объекта.

var db = utils.Db()

func ClearDb(){
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    err := db.Collection("state").Drop(ctx)
    if err != nil {
        log.Fatal(err)
    }
}
// type State struct {
//   ID        primitive.ObjectID  `json:"_id" bson:"_id"`
//   Header    string
//   Content   string
//   CreatedAt int
//   Answers   []Answer
// }
//
// type Answer struct{
//   ID        primitive.ObjectID  `json:"_id" bson:"_id"`
//   NextState primitive.ObjectID  `json:"_id" bson:"_id"`
//   Content   string
// }

func SetupStateCollection(){
    collection := db.Collection("state")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    res, err := collection.InsertOne(ctx, bson.D{
        {"Header", "Sample1"},
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
    id := res.InsertedID
    log.Println(id)

    if err != nil {
        log.Fatal(err)
    }
}
