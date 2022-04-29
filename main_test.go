package main

import (
    "telegrambot/config"
	router "telegrambot/router"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	app := InitApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/healthcheck", nil)
	app.ServeHTTP(w, req)

	response := &router.Healthcheck{
		Status: "healthy",
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, response.String(), w.Body.String())
}
