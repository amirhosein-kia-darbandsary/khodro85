package validations

import (
	"log"
	"regexp"

	"github.com/amirhosein-kia-darbandsary/khodro85/constants"
	"github.com/go-playground/validator/v10"
)

func ValidateIranianMobileNumber(fid validator.FieldLevel) bool {
	value, ok := fid.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(constants.MOBILEVALIDATION, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
