package state

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type State struct {
  ID        primitive.ObjectID  `json:"_id" bson:"_id"`
  Header    string
  Content   string
  CreatedAt int
  Answers   []Answer
}

type Answer struct{
  ID        primitive.ObjectID  `json:"_id" bson:"_id"`
  NextState primitive.ObjectID  `json:"_id" bson:"_id"`
  Content   string
}
