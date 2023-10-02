// Code generated by goa v3.13.2, DO NOT EDIT.
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
	// Multiply implements multiply.
	Multiply(context.Context, *MultiplyPayload) (res int, err error)
	// Divide implements divide.
	Divide(context.Context, *DividePayload) (res int, err error)
	// Redirect implements redirect.
	Redirect(context.Context) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"multiply", "divide", "redirect"}

// DividePayload is the payload type of the calc service divide method.
type DividePayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// MultiplyPayload is the payload type of the calc service multiply method.
type MultiplyPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// MakeZeroDivision builds a goa.ServiceError from an error.
func MakeZeroDivision(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "zero_division", false, false, false)
}
