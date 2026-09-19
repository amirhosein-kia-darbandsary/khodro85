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

// LoginUser godoc
// @Summary      Login by User Name
// @Description  Login to the service by LoginByUserNameRequest
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginByUserNameRequest true "login request"
// @Success      202 {object} base.BaseHttpResponse
// @Failure      400 {object} base.BaseHttpResponse
// @Failure      500 {object} base.BaseHttpResponse
// @Router       /user/login-by-username/ [post]
func (h *UserHandler) LoginByUserName(ctx *gin.Context) {
	req := new(dto.LoginByUserNameRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.GenerateBaseResponseWithValidationError(
			nil, false, http.StatusBadRequest, err,
		))
		return
	}
	token, err := h.UserService.LoginByUserName(*req)
	ctx.JSON(http.StatusOK, base.GenerateBaseResponse(token, true, http.StatusOK))
}

// RegisterUser godoc
// @Summary      Register by User Name
// @Description  Login to the service by RegisterRequestByUsername
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterRequestByUsername true "login request"
// @Success      202 {object} base.BaseHttpResponse
// @Failure      400 {object} base.BaseHttpResponse
// @Failure      500 {object} base.BaseHttpResponse
// @Router       /user/register-by-username/ [post]
func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	req := new(dto.RegisterRequestByUsername)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.GenerateBaseResponseWithValidationError(
			nil, false, http.StatusBadRequest, err,
		))
		return
	}
	err = h.UserService.RegisterUser(*req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			base.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err))
	}

	ctx.JSON(http.StatusOK, "Register has completed Login again")
}

// RegisterUser godoc
// @Summary      Register/Login by User Name
// @Description  Login to the service by LoginByMobileNumberRequest
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginByMobileNumberRequest true "login request"
// @Success      202 {object} base.BaseHttpResponse
// @Failure      400 {object} base.BaseHttpResponse
// @Failure      500 {object} base.BaseHttpResponse
// @Router       /user/login-by-mobile/ [post]
func (h *UserHandler) LoginOrRegisterUser(ctx *gin.Context) {
	req := new(dto.LoginByMobileNumberRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.GenerateBaseResponseWithValidationError(
			nil, false, http.StatusBadRequest, err,
		))
		return
	}

	token, err := h.UserService.RegisterLoginByMobileNumber(*req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			base.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err))
	}
	ctx.JSON(http.StatusOK, base.GenerateBaseResponse(token, true, http.StatusOK))
}
