package tests

import (
    "telegrambot/state"
)

db := Db()

func ClearDb(){
    err = db.Collection("state").Drop()
}
