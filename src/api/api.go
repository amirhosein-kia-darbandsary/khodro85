package api

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	v1 := router.Group("/api/v1/")
	{
		health_router := v1.Group("health")
		routers.Health(health_router)

	}

	router.Run(":9090")
}
