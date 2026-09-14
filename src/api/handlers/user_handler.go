package handlers

import (
	"net/http"

	"github.com/amirhosein-kia-darbandsary/khodro85/api/base"
	"github.com/amirhosein-kia-darbandsary/khodro85/api/dto"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	service := services.NewUserService(cfg)
	return &UserHandler{
		UserService: &service,
	}
}

// SendOtp godoc
// @Summary      Send OTP
// @Description  Send OTP code to the user's mobile number
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body dto.GetOtpRequest true "OTP request"
// @Success      202 {object} base.BaseHttpResponse
// @Failure      400 {object} base.BaseHttpResponse
// @Failure      500 {object} base.BaseHttpResponse
// @Router       /user/send-otp/ [post]
func (h *UserHandler) SendOtp(ctx *gin.Context) {
	otp := dto.GetOtpRequest{}

	err := ctx.ShouldBindJSON(&otp)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			base.GenerateBaseResponseWithValidationError(nil, false, 400, err),
		)
		return
	}

	err = h.UserService.SendOtp(otp)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			base.GenerateBaseResponseWithError(nil, false, 500, err),
		)
		return
	}

	ctx.JSON(
		http.StatusAccepted,
		base.GenerateBaseResponse(otp, true, 200),
	)
}
