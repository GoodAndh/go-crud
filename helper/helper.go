package helper

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// false if error is available ,otherwise true
// given struct must be pointer
func ValidateCustomStruct[p any](validate *validator.Validate, params p) (bool, map[any]any) {
	if err := validate.Struct(params); err != nil {
		errors := err.(validator.ValidationErrors)
		errorList := make(map[any]any)
		for _, er := range errors {
			var errMsg string
			va := reflect.TypeOf(params).Elem()
			field, _ := va.FieldByName(er.StructField())
			fieldName := field.Tag.Get("json")
			switch er.Tag() {
			case "required":
				errMsg = fmt.Sprintln("form ini wajib diisi")
			case "email":
				errMsg = fmt.Sprintln("bukan format email yang benar")
			case "min":
				errMsg = fmt.Sprintf(" minimal %s \n", er.Param())
			case "max":
				errMsg = fmt.Sprintf(" maximal %s \n", er.Param())
			case "eqfield":
				errMsg = fmt.Sprintf("konfirmasi password pastikan sama dengan %v \n", er.Param())
			case "number":
				errMsg = fmt.Sprintln("harus berupa angka")
			}
			errorList["error"+fieldName] = errMsg
		}
		return false, errorList
	}
	return true, map[any]any{}
}
