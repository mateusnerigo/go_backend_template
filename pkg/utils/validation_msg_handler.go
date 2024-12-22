package utils

import (
	"backend/pkg/constants"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidationMsgHandler(err error, obj interface{}) map[string]string {
	var errorsMap = make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		objType := reflect.TypeOf(obj)
		for _, fieldErr := range validationErrors {
			field, _ := objType.FieldByName(fieldErr.Field())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = fieldErr.Field()
			}

			switch fieldErr.Tag() {
			case "required":
				errorsMap[jsonTag] = constants.REQUIRED_FIELD_MISSING
			case "min":
				errorsMap[jsonTag] = constants.INVALID_MIN_FIELD_LENGTH + fieldErr.Param()
			case "max":
				errorsMap[jsonTag] = constants.INVALID_MAX_FIELD_LENGTH + fieldErr.Param()
			case "email":
				errorsMap[jsonTag] = constants.INVALID_EMAIL_FORMAT
			default:
				errorsMap[jsonTag] = "Invalid value."
			}
		}
	}
	return errorsMap
}
