package routers

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.GetHealth)
	r.POST("/", handler.PostHealth)
}
