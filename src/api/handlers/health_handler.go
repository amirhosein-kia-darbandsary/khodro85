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

// GetHealth godoc
// @Summary      Health handler
// @Description  Health API endpoint
// @Tags         Health
// @Produce      json
// @Success      200 {object} interface{}
// @Router       /health/ [get]
func (h *Health) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, base.GenerateBaseResponse("worked", true, 200))
	return
}
