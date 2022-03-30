package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Default Status = iota
	AddingNote
)

type Status int

type Handler func(*tgbotapi.BotAPI, tgbotapi.Update) Status

var userSteps = map[Status]Handler{
	Default:    HandleDefault,
	AddingNote: HandleAddingNote,
}

type UserStatuses struct {
	userStatus map[int64]Status
}

func (h *UserStatuses) GetCurrentUserHandler(userId int64) Handler {
	if _, ok := h.userStatus[userId]; !ok {
		h.userStatus[userId] = Default
	}
	return userSteps[h.userStatus[userId]]
}
func (h *UserStatuses) SetUserStatus(userId int64, status Status) {
	h.userStatus[userId] = status
}
func (h *UserStatuses) InitMap() {
	h.userStatus = make(map[int64]Status)
}
