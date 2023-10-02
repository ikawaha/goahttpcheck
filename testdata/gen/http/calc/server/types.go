// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"

	goa "goa.design/goa/v3/pkg"
)

// DivideZeroDivisionResponseBody is the type of the "calc" service "divide"
// endpoint HTTP response body for the "zero_division" error.
type DivideZeroDivisionResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewDivideZeroDivisionResponseBody builds the HTTP response body from the
// result of the "divide" endpoint of the "calc" service.
func NewDivideZeroDivisionResponseBody(res *goa.ServiceError) *DivideZeroDivisionResponseBody {
	body := &DivideZeroDivisionResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewMultiplyPayload builds a calc service multiply endpoint payload.
func NewMultiplyPayload(a int, b int) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{}
	v.A = a
	v.B = b

	return v
}

// NewDividePayload builds a calc service divide endpoint payload.
func NewDividePayload(a int, b int) *calc.DividePayload {
	v := &calc.DividePayload{}
	v.A = a
	v.B = b

	return v
}
