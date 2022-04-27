package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ChatId       string
	CurrentState primitive.ObjectID
}
