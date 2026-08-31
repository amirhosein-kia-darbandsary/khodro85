package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
}

func NewHealthHandler() *Health {
	return &Health{}
}

func (h *Health) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "worked")
}
func (h *Health) PostHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "worked")
}
