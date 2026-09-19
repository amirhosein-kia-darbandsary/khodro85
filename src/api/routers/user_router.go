package routers

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/handlers"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUserHandler(cfg)
	router.POST("/send-otp", handler.SendOtp)
	router.POST("/login-by-username", handler.LoginByUserName)
	router.POST("/register-by-username", handler.RegisterUser)
	router.POST("/login-by-mobile", handler.LoginOrRegisterUser)
}
