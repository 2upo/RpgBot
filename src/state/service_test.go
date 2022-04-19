package state

import (
	"testing"
	"telegrambot/tests"
	"github.com/stretchr/testify/assert"
)



func TestHealthcheck(t *testing.T) {

	tests.ClearDb()
	tests.SetupStateCollection()

	assert.Equal(t, 1, 1)
}
