package user

import (
	"context"
	"telegrambot/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = utils.Db().Collection("user")

func GetByChatId(chatId string) (*User, error) {
	var state User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"ChatId": chatId}).Decode(&state)

	return &state, err
}

func Upsert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(true)
	_, err := collection.ReplaceOne(ctx, bson.D{{"ChatId", user.ChatId}}, user, opts)

	return err
}
