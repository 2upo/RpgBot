package state

import (
  "telegrambot/utils"
)

type State struct {
  Id utils.Status
  Header string
  Image string
  Content string
  Answers []Answer
}

type Answer struct{
  Id string
  NextState utils.Status
  Content string
}