package api

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/middlewares"
	"github.com/amirhosein-kia-darbandsary/khodro85/api/routers"
	"github.com/amirhosein-kia-darbandsary/khodro85/api/validations"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()

	router := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("iranian_mobile", validations.ValidateIranianMobileNumber)
		val.RegisterValidation("password", validations.PasswordValidator)
	}

	router.Use(gin.Logger(), gin.Recovery(), middlewares.RateLimitter())

	v1 := router.Group("/api/v1/")
	{
		health_router := v1.Group("health")
		test_router := v1.Group("test")
		routers.Health(health_router)
		routers.Test(test_router)

	}

	router.Run(":" + cfg.PORT)
}
