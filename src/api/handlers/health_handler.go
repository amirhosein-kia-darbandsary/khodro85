package handlers

import (
	"net/http"

	"github.com/amirhosein-kia-darbandsary/khodro85/api/base"
	"github.com/gin-gonic/gin"
)

type Health struct {
}

func NewHealthHandler() *Health {
	return &Health{}
}

func (h *Health) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, base.GenerateBaseResponse("worked", true, 200))
	return
}
