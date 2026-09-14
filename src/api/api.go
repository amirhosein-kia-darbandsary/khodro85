package api

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/middlewares"
	"github.com/amirhosein-kia-darbandsary/khodro85/api/routers"
	"github.com/amirhosein-kia-darbandsary/khodro85/api/validations"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	router := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("iranian_mobile", validations.ValidateIranianMobileNumber)
		val.RegisterValidation("password", validations.PasswordValidator)
	}

	router.Use(gin.Logger(), gin.Recovery(), middlewares.RateLimitter(), middlewares.DefaultStructuredLogger(cfg))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1/")
	{
		health_router := v1.Group("health")
		test_router := v1.Group("test")
		user_router := v1.Group("user")
		routers.Health(health_router)
		routers.Test(test_router)
		routers.User(user_router, cfg)

	}

	router.Run(":" + cfg.PORT)
}
