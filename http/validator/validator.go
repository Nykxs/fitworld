package httpvalidator

import (
	"github.com/asaskevich/govalidator"
)

// Validator implements our validator domain interface.
// It is used by each http endpoint to validate received parameters.
// Improvements : the Validator struct can be improved with fields to have a better error handling.
type Validator struct {
}

// NewValidator returns a Validator that can be used by any http endpoint to validate its parameters.
func NewValidator() *Validator {
	return &Validator{}
}

// Validate takes a struct as parameter and validates it using fieldtag.
// fields without field tags are ignored.
func (v *Validator) Validate(payload interface{}) error {
	ok, err := govalidator.ValidateStruct(payload)
	if ok {
		return nil
	}
	return err
}
