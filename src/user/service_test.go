package user

import (
	"context"
	"telegrambot/tests"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userService = InitUserService()

func TestGetByChatId(t *testing.T) {
	ass := assert.New(t)
	chatId := "aboba"
	_, err := tests.NewUser(chatId, userService.Collection)
	ass.Nil(err)

	state, err := userService.GetByChatId(chatId)

	ass.Nil(err)
	ass.Equal(state.ChatId, chatId)

	tests.ClearDb([]*mongo.Collection{userService.Collection}) // Tear down
}

func TestUpsert(t *testing.T) {
	ass := assert.New(t)
	chatId := "aboba"
	user := User{
		ChatId:       chatId,
		CurrentState: primitive.NewObjectID(),
	}
	err := userService.Upsert(&user)
	ass.Nil(err)

	var createdUser User
	err = userService.Collection.FindOne(context.Background(), bson.D{{"chatid", chatId}}).Decode(&createdUser)
	ass.Nil(err)
	ass.Equal(createdUser.ChatId, user.ChatId)
	ass.Equal(createdUser.CurrentState, user.CurrentState)

}

// func TestUpdateUserState(t *testing.T) {

// }
