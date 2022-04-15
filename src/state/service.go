package state

import (
    "log"
    "telegrambot/utils"
)

db = Db()

func GetAll() []State {
  states := make([]State)
  
  ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()
  cursor, err := db.Find(ctx, bson.D{})
  
  if err != nil { log.Fatal(err) }
  
  defer cursor.Close(ctx)

  for cursor.Next(ctx) {
      var result State
      err := cursor.Decode(&result)
      if err != nil { log.Fatal(err) }
      states = append(states, result)
  }
  
  if err := cursor.Err(); err != nil {
      log.Fatal(err)
  }
}

func Insert(new_state State) State{
  
}

func GetById(id utils.Status) State {
  
}

func DeleteById(id utils.Status) {
  
}

func Update(updated_state State) State{
  
}


