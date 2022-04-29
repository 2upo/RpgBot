package router

import (
	"net/http"
    "encoding/json"
	"github.com/gin-gonic/gin"
)


func healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

type Healthcheck struct {
	Status string `json:"status"`
}

// Serialize Healthcheck response to string
func (h *Healthcheck)String() string {
    serialized_response, _ := json.Marshal(h)

    return string(serialized_response)
}
