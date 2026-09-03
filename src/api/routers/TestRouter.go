package routers

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/handlers"
	"github.com/gin-gonic/gin"
)

func Test(router *gin.RouterGroup) {
	handlers := handlers.NewTest()

	router.GET("/", handlers.TestHandler)
	router.POST("/", handlers.TestBindingHandler)
}
