package user

import (
  "telegrambot/utils"
)

var db = utils.Db()

// Get user state from MongoDB
func GetUserState(userId string) utils.Status {
    var result struct {
        Value float64
    }
    filter := bson.D{{"name", "pi"}}
    ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err = collection.FindOne(ctx, filter).Decode(&result)
    if err == mongo.ErrNoDocuments {
        // Do something when no record was found
        fmt.Println("record does not exist")
    } else if err != nil {
        log.Fatal(err)
    }
    // Do something with result...
}

// Set new User Status on DB
func SetUserState(userId string, newState utils.Status) {

}
