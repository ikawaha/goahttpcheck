// Code generated by goa v3.0.9, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The calc service performs operations on numbers.
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
	// Div implements div.
	Div(context.Context, *DivPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"add", "div"}

// AddPayload is the payload type of the calc service add method.
type AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// DivPayload is the payload type of the calc service div method.
type DivPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// MakeZeroDivision builds a goa.ServiceError from an error.
func MakeZeroDivision(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "zero_division",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
