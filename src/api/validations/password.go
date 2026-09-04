package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	upperRegex   = regexp.MustCompile(`[A-Z]`)
	lowerRegex   = regexp.MustCompile(`[a-z]`)
	numberRegex  = regexp.MustCompile(`[0-9]`)
	specialRegex = regexp.MustCompile(`[!@#$%^&*]`)
	lengthRegex  = regexp.MustCompile(`^.{8,}$`)
)

func PasswordValidator(fid validator.FieldLevel) bool {
	val, ok := fid.Field().Interface().(string)
	if !ok {
		return false
	}

	return upperRegex.MatchString(val) &&
		lowerRegex.MatchString(val) &&
		numberRegex.MatchString(val) &&
		specialRegex.MatchString(val) &&
		lengthRegex.MatchString(val)
}
