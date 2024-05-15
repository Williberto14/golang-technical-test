package utils

import "github.com/go-playground/validator/v10"

// GetValidator returns a new instance of the validator.Validate struct.
func GetValidator() *validator.Validate {
	return validator.New()
}
