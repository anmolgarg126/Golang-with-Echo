package utility

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// var validator *validator.Validate = validator.New()
var v = validator.New()

func ValidateStruct(i interface{}) error {
	fmt.Println("Validating struct :: ", i)
	err := v.Struct(i)
	fmt.Println("err while validating :: ", err)
	return err
}
