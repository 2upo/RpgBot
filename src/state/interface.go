package state

import (
	"context"
	"telegrambot/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IStateService interface {
	GetAll() ([]State, error)
    Insert(new_state *State) error
    GetById(id primitive.ObjectID) (*State, error)
    Update(updated_state *State) error
    DeleteById(id primitive.ObjectID) error
}
