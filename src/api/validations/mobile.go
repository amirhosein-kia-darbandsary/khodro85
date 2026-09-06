package validations

import (
	"regexp"

	"github.com/amirhosein-kia-darbandsary/khodro85/constants"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"github.com/go-playground/validator/v10"
)

func ValidateIranianMobileNumber(fid validator.FieldLevel) bool {
	value, ok := fid.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(constants.MOBILEVALIDATION, value)
	if err != nil {
		logger.Error(logging.Category(logging.MobileValidation), logging.MobileValidation, err.Error(), nil)
	}
	return res
}
