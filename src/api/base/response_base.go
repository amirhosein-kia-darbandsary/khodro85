package base

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/validations"
)

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"resultcode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultcode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultcode,
		ValidationErrors: nil,
		Error:            nil,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultcode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultcode,
		ValidationErrors: nil,
		Error:            err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resulcode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resulcode,
		ValidationErrors: validations.ErrorValidation(err),
		Error:            err.Error(),
	}
}
