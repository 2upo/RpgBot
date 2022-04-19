package state

import (
    "log"
    "context"
    "telegrambot/utils"
    "go.mongodb.org/mongo-driver/mongo/options"
)

db = Db()

// Get all states from database
func GetAll() []State {
  collection = db.Collection("state")
  states := make([]State)

  ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  // defining select options
  options := options.Find()
  options.SetSort(bson.D{{"CreatedAt", -1}})

  // find all
  // TODO: add pagination
  // options.SetLimit(10)
  // options.SetSkip(10)

  cursor, err := collection.Find(ctx, bson.D{}})
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
  return result, nil
}
//
// func Insert(new_state State) State {
// }
//
// func GetById(id utils.Status) State {
//
// }
//
// func DeleteById(id utils.Status) {
//
// }
//
// func Update(updated_state State) State {
//
// }
