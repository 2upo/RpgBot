package user

import (
	"context"
	"telegrambot/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	Collection *mongo.Collection
}

func InitUserService() *UserService {
	var userService UserService
	userService.Collection = utils.Db().Collection("user")

	return &userService
}

func (service *UserService) GetByChatId(chatId string) (*User, error) {
	var state User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := service.Collection.FindOne(ctx, bson.M{"chatid": chatId}).Decode(&state)

	return &state, err
}

func (service *UserService) Upsert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := options.Replace().SetUpsert(true)
	_, err := service.Collection.ReplaceOne(ctx, bson.D{{"chatid", user.ChatId}}, user, opts)

	return err
}
