package utils

import (
    ""
)

db = Db()

func (h *UserStatuses) GetCurrentUserHandler(userId int64) Handler {
	if _, ok := h.userStatus[userId]; !ok {
		h.userStatus[userId] = Default
	}
	return userSteps[h.userStatus[userId]]
}
func SetUserStatus(userId int64, status Status) {
	h.userStatus[userId] = status
}
