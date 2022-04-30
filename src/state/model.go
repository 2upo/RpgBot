package state

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type State struct {
  ID        primitive.ObjectID  `json:"_id"         bson:"_id"        binding:"isdefault"`
  Header    string              `json:"user"        bson:"header"     binding:"required,alphanum"`
  Content   string              `json:"content"     bson:"content"    binding:"max=500"`
  CreatedAt int                 `json:"createdat"   bson:"createdat"  binding:"isdefault"`
  Answers   []Answer            `json:"answers"     bson:"answers"`
}

type Answer struct{
  NextState primitive.ObjectID  `json:"nextstate" bson:"nextstate"`
  Content   string              `json:"content"   bson:"content"      binding:"max=500,required_with=nextstate"`
}
