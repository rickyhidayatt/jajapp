package input

import (
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	Validator *validator.Validate
}

var validation Validation
var validationSingleton sync.Once

func GetValidator() Validation {
	validationSingleton.Do(func() {
		validator := validator.New()
		validator.RegisterValidation("valid_address", validateAlphaNumericWithSpaceDot, true)
		validator.RegisterValidation("valid_string", validateAlphanumericWithSpace, true)
		validator.RegisterValidation("valid_nik", validateNik, true)
		validator.RegisterValidation("valid_phone_number", phoneNumberValidator, true)

		validation = Validation{
			Validator: validator,
		}
	})

	return validation
}

func validateNik(fl validator.FieldLevel) bool {
	nikRegex := regexp.MustCompile(`^[0-9]{16}$`)
	return nikRegex.MatchString(fl.Field().String())
}

func validateAlphaNumericWithSpaceDot(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9\\s.]+$")
	return regex.MatchString(fl.Field().String())
}

func validateAlphanumericWithSpace(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[a-zA-Z\\s]+$")
	return regex.MatchString(fl.Field().String())
}

func phoneNumberValidator(fl validator.FieldLevel) bool {
	phoneRegex := regexp.MustCompile(`^[0-9]{0,13}$`)
	return phoneRegex.MatchString(fl.Field().String())
}

func ValidateUserRequest(params interface{}) error {
	validate := GetValidator().Validator
	validate.RegisterValidation("name_validator", validateAlphanumericWithSpace, true)
	validate.RegisterValidation("phone_number_validator", phoneNumberValidator, true)
	validate.RegisterValidation("address_validator", validateAlphaNumericWithSpaceDot, true)
	validate.RegisterValidation("nik_validator", validateNik, true)

	return validate.Struct(params)
}
