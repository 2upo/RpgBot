package state

import (
  "telegrambot/utils"
)

type State struct {
  Id utils.Status
  Header string
  Content string
  CreatedAt int
  Answers []Answer
}

type Answer struct{
  Id string
  NextState utils.Status
  Content string
}
