package state

import (
    "log"
    "context"
    "telegrambot/utils"
    "time"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = utils.Db().Collection("state")

// Get all states from database
func GetAll() ([]State, error) {
  var states []State

  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  // defining select options
  options := options.Find()
  options.SetSort(bson.D{{"CreatedAt", -1}})

  // find all
  // TODO: add pagination
  // options.SetLimit(10)
  // options.SetSkip(10)

  cursor, err := collection.Find(ctx, bson.D{}, options)
  if err != nil {
      log.Fatal(err)
      return nil, err
  }

  defer cursor.Close(ctx)

  for cursor.Next(ctx) {
      var result State

      err := cursor.Decode(&result)
      if err != nil {
          log.Fatal(err)
          return nil, err
      }
      states = append(states, result)
  }

  if err := cursor.Err(); err != nil {
      log.Fatal(err)
      return nil, err
  }
  return states, nil
}
//
// func Insert(new_state State) State {
// }
//
func GetById(id primitive.ObjectID) (*State, error) {
    var state State

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // defining select options
    options := options.FindOne()
    options.SetSort(bson.D{{"CreatedAt", -1}})

    // find all
    // TODO: add pagination
    // options.SetLimit(10)
    // options.SetSkip(10)

    err := collection.FindOne(ctx, bson.M{"_id": id}, options).Decode(&state)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return &state, nil
}
//
// func DeleteById(id utils.Status) {
//
// }
//
// func Update(updated_state State) State {
//
// }
//
